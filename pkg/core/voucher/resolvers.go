package voucher

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

type VouchersResolver struct {
	VouchersRepo *coredb.VouchersRepo
}

func NewVouchersResolver() *VouchersResolver {
	return &VouchersResolver{
		VouchersRepo: coredb.NewVouchersRepo(),
	}
}

func (r *VouchersResolver) CreateVoucher(params graphql.ResolveParams) (interface{}, error) {
	voucher := coredb.Voucher{
		ID:          primitive.NewObjectID(),
		Code:        params.Args["code"].(string),
		BrandId:     params.Args["brandId"].(string),
		ImageURL:    params.Args["imageURL"].(string),
		Value:       params.Args["value"].(string),
		Description: params.Args["description"].(string),
		ExpiredDate: params.Args["expiredDate"].(time.Time),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := db.GetVoucherCollection().InsertOne(ctx, voucher)
	if err != nil {
		log.Printf("failed to insert voucher: %v\n", err)
		return false, err
	}

	return true, nil
}

func (r *VouchersResolver) GetVoucherByID(params graphql.ResolveParams) (interface{}, error) {
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

	var voucher coredb.Voucher
	err = db.GetVoucherCollection().FindOne(ctx, bson.M{"_id": objectID}).Decode(&voucher)
	if err != nil {
		log.Printf("failed to find voucher: %v\n", err)
		return nil, err
	}

	return voucher, nil
}

func (r *VouchersResolver) GetVoucherByCode(params graphql.ResolveParams) (interface{}, error) {
	code, ok := params.Args["code"].(string)
	if !ok {
		return nil, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var voucher coredb.Voucher
	err := db.GetVoucherCollection().FindOne(ctx, bson.M{"code": code}).Decode(&voucher)
	if err != nil {
		log.Printf("failed to find voucher: %v\n", err)
		return nil, err
	}

	return voucher, nil
}

func (r *VouchersResolver) GetAllVouchers(params graphql.ResolveParams) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var vouchers []coredb.Voucher
	cursor, err := db.GetVoucherCollection().Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to find vouchers: %v", err)
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &vouchers); err != nil {
		return nil, fmt.Errorf("failed to decode vouchers: %v", err)
	}

	return vouchers, nil
}

func (r *VouchersResolver) GetVouchersByBrandId(params graphql.ResolveParams) (interface{}, error) {
	brandId, ok := params.Args["brandId"].(string)
	if !ok {
		return nil, fmt.Errorf("missing brand ID")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var vouchers []coredb.Voucher
	cursor, err := db.GetVoucherCollection().Find(ctx, bson.M{"brand_id": brandId})
	if err != nil {
		return nil, fmt.Errorf("failed to find vouchers: %v", err)
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &vouchers); err != nil {
		return nil, fmt.Errorf("failed to decode vouchers: %v", err)
	}

	return vouchers, nil
}

func (r *VouchersResolver) EditVoucher(params graphql.ResolveParams) (interface{}, error) {
	user, ok := params.Context.Value(auth.UserKey).(coredb.User)
	if !ok {
		return nil, fmt.Errorf("user not found")
	}

	if user.Role != "brand" {
		return nil, fmt.Errorf("Permission denied")
	}

	id, ok := params.Args["id"].(string)
	if !ok {
		return nil, fmt.Errorf("missing voucher ID")
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid voucher ID")
	}

	brandId, ok := params.Args["brandId"].(string)
	if !ok {
		return nil, fmt.Errorf("missing brand ID")
	}

	brandObjectID, err := primitive.ObjectIDFromHex(brandId)
	if err != nil {
		return nil, fmt.Errorf("invalid brand ID")
	}

	if user.ID != brandObjectID {
		return nil, fmt.Errorf("Brand ID does not match")
	}

	update := bson.M{}
	if code, ok := params.Args["code"].(string); ok {
		update["code"] = code
	}
	if imageURL, ok := params.Args["imageURL"].(string); ok {
		update["imageURL"] = imageURL
	}
	if value, ok := params.Args["value"].(string); ok {
		update["value"] = value
	}
	if description, ok := params.Args["description"].(string); ok {
		update["description"] = description
	}
	if expiredDate, ok := params.Args["expiredDate"].(time.Time); ok {
		update["expiredDate"] = expiredDate
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := db.GetVoucherCollection().UpdateOne(
		ctx,
		bson.M{"_id": objectID, "brandId": brandId},
		bson.M{"$set": update},
	)
	if err != nil {
		log.Printf("failed to update voucher: %v\n", err)
		return nil, fmt.Errorf("failed to update voucher: %v", err)
	}

	if result.MatchedCount == 0 {
		return nil, fmt.Errorf("no matching voucher found")
	}

	return true, nil
}

func (r *VouchersResolver) DeleteVoucher(params graphql.ResolveParams) (interface{}, error) {
	user, ok := params.Context.Value(auth.UserKey).(coredb.User)
	if !ok {
		return nil, fmt.Errorf("user not found")
	}

	if user.Role != "brand" {
		return nil, fmt.Errorf("Permission denied")
	}

	id, ok := params.Args["id"].(string)
	if !ok {
		return nil, fmt.Errorf("missing voucher ID")
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid voucher ID")
	}

	brandId, ok := params.Args["brandId"].(string)
	if !ok {
		return nil, fmt.Errorf("missing brand ID")
	}

	brandObjectID, err := primitive.ObjectIDFromHex(brandId)
	if err != nil {
		return nil, fmt.Errorf("invalid brand ID")
	}

	if user.ID != brandObjectID {
		return nil, fmt.Errorf("Brand ID does not match")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := db.GetVoucherCollection().DeleteOne(ctx, bson.M{"_id": objectID, "brandId": brandObjectID})
	if err != nil {
		log.Printf("failed to delete voucher: %v\n", err)
		return nil, fmt.Errorf("failed to delete voucher: %v", err)
	}

	if result.DeletedCount == 0 {
		return nil, fmt.Errorf("no matching voucher found or you don't have permission to delete it")
	}

	return true, nil
}

func (r *VouchersResolver) GetVouchersByUserID(params graphql.ResolveParams) (interface{}, error) {
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
			return []coredb.Voucher{}, nil
		}
		log.Printf("failed to fetch package: %v\n", err)
		return nil, err
	}

	log.Printf("pkg.Vouchers: %v", pkg.Vouchers)

	var vouchers []coredb.Voucher

	for _, voucherIDStr := range pkg.Vouchers {
		objectID, err := primitive.ObjectIDFromHex(voucherIDStr)
		if err != nil {
			log.Printf("invalid voucher ID: %v\n", err)
			return nil, err
		}

		var voucher coredb.Voucher
		err = db.GetVoucherCollection().FindOne(ctx, bson.M{"_id": objectID}).Decode(&voucher)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				continue
			}
			log.Printf("failed to fetch voucher: %v\n", err)
			return nil, err
		}

		vouchers = append(vouchers, voucher)
	}

	log.Printf("Fetched vouchers: %v", vouchers)

	return vouchers, nil
}
