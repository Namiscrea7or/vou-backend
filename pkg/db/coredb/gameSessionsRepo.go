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

type GameSessionsRepo struct {
	*mongo.Collection
}

func NewGameSessionsRepo() *GameSessionsRepo {
	gameSessionsCollection := db.GetGameSessionsCollection()
	_, err := gameSessionsCollection.Indexes().CreateMany(context.TODO(), []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "start_time", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	})
	if err != nil {
		log.Println("failed to create indexes for game sessions collection")
		return nil
	}

	return &GameSessionsRepo{gameSessionsCollection}
}

func (r *GameSessionsRepo) GetGameSessionByID(id primitive.ObjectID) (GameSession, error) {
	var gameSession GameSession
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.FindOne(ctx, bson.M{"_id": id}).Decode(&gameSession)

	return gameSession, err
}

func (r *GameSessionsRepo) CreateNewGameSession(gameSession GameSession) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return r.InsertOne(ctx, gameSession)
}
