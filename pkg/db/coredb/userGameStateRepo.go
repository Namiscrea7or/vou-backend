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

type UserGameStatesRepo struct {
	*mongo.Collection
}

func NewUserGameStatesRepo() *UserGameStatesRepo {
	userGameStatesCollection := db.GetUserGameStatesCollection()
	_, err := userGameStatesCollection.Indexes().CreateMany(context.TODO(), []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "user_id", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	})
	if err != nil {
		log.Println("failed to create indexes for user game states collection")
		return nil
	}

	return &UserGameStatesRepo{userGameStatesCollection}
}

func (r *UserGameStatesRepo) GetUserGameStateByUserID(userID primitive.ObjectID) (UserGameState, error) {
	var userGameState UserGameState
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.FindOne(ctx, bson.M{"user_id": userID}).Decode(&userGameState)

	return userGameState, err
}

func (r *UserGameStatesRepo) UpdateUserGameState(userGameState UserGameState) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"user_id": userGameState.UserID}
	update := bson.M{
		"$set": userGameState,
	}

	return r.UpdateOne(ctx, filter, update)
}
