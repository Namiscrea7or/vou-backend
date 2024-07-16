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

type BrandRepo struct {
	*mongo.Collection
}

func NewBrandRepo() *BrandRepo {
	brandCollection := db.GetBrandCollection()
	_, err := brandCollection.Indexes().CreateMany(context.TODO(), []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "creatorId", Value: 1}},
			Options: options.Index().SetUnique(false),
		},
	})
	if err != nil {
		log.Println("failed to create indexes for brands collection")
		return nil
	}

	return &BrandRepo{brandCollection}
}

func (r *BrandRepo) GetAllBrands() ([]Brand, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var brands []Brand
	cursor, err := r.Find(ctx, primitive.M{})
	if err != nil {
		log.Printf("failed to fetch Brands: %v\n", err)
		return nil, err
	}

	defer cursor.Close(ctx)

	err = cursor.All(ctx, &brands)
	if err != nil {
		log.Printf("failed to decode brands: %v\n", err)
		return nil, err
	}

	return brands, nil
}

func (r *BrandRepo) GetBrandByID(id primitive.ObjectID) (Brand, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	brand := Brand{}
	err := r.FindOne(ctx, bson.M{"_id": id}).Decode(&brand)

	return brand, err
}

func (r *BrandRepo) CreateBrand(brand Brand) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := r.InsertOne(ctx, brand)

	return result, err
}
