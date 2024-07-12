package coredb

import (
	"context"
	"log"
	"time"

	"vou/pkg/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type VouchersRepo struct {
	*mongo.Collection
}

func NewVouchersRepo() *VouchersRepo {
	vouchersCollection := db.GetVoucherCollection()
	_, err := vouchersCollection.Indexes().CreateMany(context.TODO(), []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "code", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	})
	if err != nil {
		log.Println("failed to create indexes for vouchers collection")
		return nil
	}

	return &VouchersRepo{vouchersCollection}
}

func (r *VouchersRepo) CreateNewVoucher(voucher Voucher) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return r.InsertOne(ctx, voucher)
}

func (r *VouchersRepo) GetVoucherByCode(code string) (*Voucher, error) {
	var voucher Voucher
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.Collection.FindOne(ctx, bson.M{"code": code}).Decode(&voucher)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &voucher, nil
}

func (r *VouchersRepo) GetVoucherByID(id primitive.ObjectID) (*Voucher, error) {
	var voucher Voucher
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.Collection.FindOne(ctx, bson.M{"_id": id}).Decode(&voucher)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &voucher, nil
}
