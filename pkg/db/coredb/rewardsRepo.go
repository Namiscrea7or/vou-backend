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

type RewardsRepo struct {
	*mongo.Collection
}

func NewRewardsRepo() *RewardsRepo {
	rewardsCollection := db.GetRewardsCollection()
	_, err := rewardsCollection.Indexes().CreateMany(context.TODO(), []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "name", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	})
	if err != nil {
		log.Println("failed to create indexes for rewards collection")
		return nil
	}

	return &RewardsRepo{rewardsCollection}
}

func (r *RewardsRepo) GetRewardByID(id primitive.ObjectID) (Reward, error) {
	var reward Reward
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.FindOne(ctx, bson.M{"_id": id}).Decode(&reward)

	return reward, err
}

func (r *RewardsRepo) CreateNewReward(reward Reward) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return r.InsertOne(ctx, reward)
}
