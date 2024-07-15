package voucher

import (
	"context"
	"log"
	"time"

	"vou/pkg/db"
	"vou/pkg/db/coredb"

	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		ImageURL:    params.Args["imageURL"].(string),
		Value:       params.Args["value"].(float64),
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
