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

type ExchangesRepo struct {
	*mongo.Collection
}

func NewExchangesRepo() *ExchangesRepo {
	exchangesCollection := db.GetExchangeCollection()
	_, err := exchangesCollection.Indexes().CreateMany(context.TODO(), []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "firstUserId", Value: 1}},
			Options: options.Index().SetUnique(false),
		},
		{
			Keys:    bson.D{{Key: "secondUserId", Value: 1}},
			Options: options.Index().SetUnique(false),
		},
	})
	if err != nil {
		log.Println("failed to create indexes for exchanges collection")
		return nil
	}

	return &ExchangesRepo{exchangesCollection}
}

func (r *ExchangesRepo) GetExchangeByID(id primitive.ObjectID) (Exchange, error) {
	var exchange Exchange
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.Collection.FindOne(ctx, bson.M{"_id": id}).Decode(&exchange)
	return exchange, err
}

func (r *ExchangesRepo) GetUncompletedExchanges(ctx context.Context) ([]Exchange, error) {
	var exchanges []Exchange
	cursor, err := r.Collection.Find(ctx, bson.M{"completed": false})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var exchange Exchange
		if err := cursor.Decode(&exchange); err != nil {
			return nil, err
		}
		exchanges = append(exchanges, exchange)
	}

	return exchanges, nil
}

func (r *ExchangesRepo) CreateNewExchange(exchange Exchange) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return r.Collection.InsertOne(ctx, exchange)
}

func (r *ExchangesRepo) AddVoucherToExchange(exchangeID primitive.ObjectID, secondUserID, secondVoucherCode string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": exchangeID, "completed": false}
	update := bson.M{
		"$set": bson.M{
			"secondUserId":      secondUserID,
			"secondVoucherCode": secondVoucherCode,
		},
	}

	_, err := r.Collection.UpdateOne(ctx, filter, update)
	return err
}

func (r *ExchangesRepo) UpdateExchange(ctx context.Context, exchange Exchange) error {
	filter := bson.M{"_id": exchange.ID}
	update := bson.M{"$set": exchange}
	_, err := r.Collection.UpdateOne(ctx, filter, update)
	return err
}
