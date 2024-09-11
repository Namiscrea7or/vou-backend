package packages

import (
	"context"
	"fmt"
	"log"
	"time"

	"vou/pkg/auth"
	"vou/pkg/db"
	"vou/pkg/db/coredb"

	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PackagesResolver struct {
	PackagesRepo *coredb.PackagesRepo
}

func NewPackagesResolver() *PackagesResolver {
	return &PackagesResolver{
		PackagesRepo: coredb.NewPackagesRepo(),
	}
}

func (r *PackagesResolver) CreatePackage(params graphql.ResolveParams) (interface{}, error) {
	pkg := coredb.Package{
		ID:            primitive.NewObjectID(),
		UserID:        params.Args["userId"].(string),
		Vouchers:      []string{},
		Rewards:       []string{},
		AllowExchange: params.Args["allowExchange"].(bool),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := db.GetPackageCollection().InsertOne(ctx, pkg)
	if err != nil {
		log.Printf("failed to insert package: %v\n", err)
		return false, err
	}

	return true, nil
}

func (r *PackagesResolver) GetPackageByID(params graphql.ResolveParams) (interface{}, error) {
	user, ok := params.Context.Value(auth.UserKey).(coredb.User)
	if !ok {
		return nil, fmt.Errorf("user not found")
	}

	if user.Role != "user" || user.Role != "admin" {
		return nil, fmt.Errorf("Permission denied")
	}

	id, ok := params.Args["id"].(string)
	if !ok {
		return nil, nil
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var pkg coredb.Package
	err = db.GetPackageCollection().FindOne(ctx, map[string]primitive.ObjectID{
		"_id": objectID,
	}).Decode(&pkg)
	if err != nil {
		log.Printf("failed to find package: %v\n", err)
		return nil, err
	}

	return pkg, nil
}

func (r *PackagesResolver) AddRewardToPackageById(params graphql.ResolveParams) (interface{}, error) {
	packageID, _ := params.Args["packageID"].(string)
	rewardID, _ := params.Args["rewardID"].(string)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pkgID, err := primitive.ObjectIDFromHex(packageID)
	if err != nil {
		return false, err
	}

	rwID, err := primitive.ObjectIDFromHex(rewardID)
	if err != nil {
		return false, err
	}

	filter := bson.M{"_id": pkgID}
	update := bson.M{
		"$push": bson.M{"rewards": rwID},
	}

	_, err = r.PackagesRepo.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *PackagesResolver) RemoveOneRewardFromPackageById(params graphql.ResolveParams) (interface{}, error) {
	packageID, _ := params.Args["packageID"].(string)
	rewardID, _ := params.Args["rewardID"].(string)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pkgID, err := primitive.ObjectIDFromHex(packageID)
	if err != nil {
		return false, err
	}

	rwID, err := primitive.ObjectIDFromHex(rewardID)
	if err != nil {
		return false, err
	}

	var pkg coredb.Package
	err = db.GetPackageCollection().FindOne(ctx, bson.M{"_id": pkgID}).Decode(&pkg)
	if err != nil {
		return false, err
	}

	found := false
	for i, reward := range pkg.Rewards {
		if reward == rwID.Hex() {
			pkg.Rewards = append(pkg.Rewards[:i], pkg.Rewards[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		return false, fmt.Errorf("reward not found in package")
	}

	filter := bson.M{"_id": pkgID}
	update := bson.M{"$set": bson.M{"rewards": pkg.Rewards}}

	_, err = r.PackagesRepo.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *PackagesResolver) AddVoucherToPackageByID(params graphql.ResolveParams) (interface{}, error) {
	packageID, _ := params.Args["packageID"].(string)
	voucherID, _ := params.Args["voucherID"].(string)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pkgID, err := primitive.ObjectIDFromHex(packageID)
	if err != nil {
		return false, fmt.Errorf("invalid package ID: %v", err)
	}

	vID, err := primitive.ObjectIDFromHex(voucherID)
	if err != nil {
		return false, fmt.Errorf("invalid voucher ID: %v", err)
	}

	var voucher coredb.Voucher
	err = db.GetVoucherCollection().FindOne(ctx, bson.M{"_id": vID}).Decode(&voucher)
	if err != nil {
		return false, fmt.Errorf("voucher not found: %v", err)
	}

	filter := bson.M{"_id": pkgID}
	update := bson.M{
		"$push": bson.M{"vouchers": vID.Hex()},
	}

	_, err = r.PackagesRepo.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return false, fmt.Errorf("failed to update package: %v", err)
	}

	return true, nil
}

func (r *PackagesResolver) RemoveOneVoucherFromPackageByID(params graphql.ResolveParams) (interface{}, error) {
	packageID, _ := params.Args["packageID"].(string)
	voucherID, _ := params.Args["voucherID"].(string)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pkgID, err := primitive.ObjectIDFromHex(packageID)
	if err != nil {
		return false, fmt.Errorf("invalid package ID: %v", err)
	}

	vID, err := primitive.ObjectIDFromHex(voucherID)
	if err != nil {
		return false, fmt.Errorf("invalid voucher ID: %v", err)
	}

	var pkg coredb.Package
	err = db.GetPackageCollection().FindOne(ctx, bson.M{"_id": pkgID}).Decode(&pkg)
	if err != nil {
		return false, fmt.Errorf("failed to find package: %v", err)
	}

	found := false
	for i, voucher := range pkg.Vouchers {
		if voucher == vID.Hex() {
			pkg.Vouchers = append(pkg.Vouchers[:i], pkg.Vouchers[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		return false, fmt.Errorf("voucher not found in package")
	}

	filter := bson.M{"_id": pkgID}
	update := bson.M{"$set": bson.M{"vouchers": pkg.Vouchers}}

	_, err = r.PackagesRepo.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return false, fmt.Errorf("failed to update package: %v", err)
	}

	return true, nil
}

func (r *PackagesResolver) GetPackageByUserID(params graphql.ResolveParams) (interface{}, error) {
	userID, ok := params.Args["userId"].(string)
	if !ok {
		return nil, fmt.Errorf("invalid user ID")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var pkg coredb.Package
	err := db.GetPackageCollection().FindOne(ctx, bson.M{"user_id": userID}).Decode(&pkg)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("no package found for user ID: %v", userID)
		}
		log.Printf("failed to fetch package: %v", err)
		return nil, fmt.Errorf("failed to fetch package: %v", err)
	}

	return pkg, nil
}
