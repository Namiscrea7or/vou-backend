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
		AllowExchange: params.Args["allow_exchange"].(bool),
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
	user, ok := params.Context.Value(auth.UserKey).(coredb.User)
	if !ok {
		return nil, fmt.Errorf("user not found")
	}

	if user.Role != "user" {
		return nil, fmt.Errorf("Permission denied")
	}

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
		"$addToSet": bson.M{"rewards": rwID},
	}

	_, err = r.PackagesRepo.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *PackagesResolver) RemoveRewardFromPackageById(params graphql.ResolveParams) (interface{}, error) {
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
		"$pull": bson.M{"rewards": rwID},
	}

	_, err = r.PackagesRepo.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *PackagesResolver) AddVoucherToPackageByCode(params graphql.ResolveParams) (interface{}, error) {
	user, ok := params.Context.Value(auth.UserKey).(coredb.User)
	if !ok {
		return nil, fmt.Errorf("user not found")
	}

	if user.Role != "user" {
		return nil, fmt.Errorf("Permission denied")
	}

	packageID, _ := params.Args["packageID"].(string)
	voucherCode, _ := params.Args["voucherCode"].(string)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pkgID, err := primitive.ObjectIDFromHex(packageID)
	if err != nil {
		return false, err
	}

	var voucher coredb.Voucher
	err = db.GetVoucherCollection().FindOne(ctx, bson.M{"code": voucherCode}).Decode(&voucher)
	if err != nil {
		return false, err
	}

	filter := bson.M{"_id": pkgID}
	update := bson.M{
		"$addToSet": bson.M{"vouchers": voucher.ID.Hex()},
	}

	_, err = r.PackagesRepo.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *PackagesResolver) RemoveVoucherFromPackageByCode(params graphql.ResolveParams) (interface{}, error) {
	packageID, _ := params.Args["packageID"].(string)
	voucherCode, _ := params.Args["voucherCode"].(string)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pkgID, err := primitive.ObjectIDFromHex(packageID)
	if err != nil {
		return false, err
	}

	var voucher coredb.Voucher
	err = db.GetVoucherCollection().FindOne(ctx, bson.M{"code": voucherCode}).Decode(&voucher)
	if err != nil {
		return false, err
	}

	filter := bson.M{"_id": pkgID}
	update := bson.M{
		"$pull": bson.M{"vouchers": voucher.ID.Hex()},
	}

	_, err = r.PackagesRepo.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return false, err
	}

	return true, nil
}
