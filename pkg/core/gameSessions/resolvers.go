package gameSessions

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

type GameSessionsResolver struct {
	GameSessionsRepo *coredb.GameSessionsRepo
}

func NewGameSessionsResolver() *GameSessionsResolver {
	return &GameSessionsResolver{
		GameSessionsRepo: coredb.NewGameSessionsRepo(),
	}
}

func (r *GameSessionsResolver) CreateGameSession(params graphql.ResolveParams) (interface{}, error) {
	gameSession := coredb.GameSession{
		ID:        primitive.NewObjectID(),
		StartTime: time.Now(),
		EndTime:   time.Now(), // change later
		Status:    true,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := db.GetGameSessionsCollection().InsertOne(ctx, gameSession)
	if err != nil {
		log.Printf("failed to insert game session: %v\n", err)
		return false, err
	}

	return true, nil
}

func (r *GameSessionsResolver) GetGameSessionByID(params graphql.ResolveParams) (interface{}, error) {
	id, err := primitive.ObjectIDFromHex(params.Args["id"].(string))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var gameSession coredb.GameSession
	err = db.GetGameSessionsCollection().FindOne(ctx, map[string]primitive.ObjectID{
		"_id": id,
	}).Decode(&gameSession)
	if err != nil {
		log.Printf("failed to find game session: %v\n", err)
		return nil, err
	}

	return gameSession, nil
}

func (r *GameSessionsResolver) AddRewardToGameSession(params graphql.ResolveParams) (interface{}, error) {
	ID, _ := params.Args["gameSessionID"].(string)
	gameId, _ := primitive.ObjectIDFromHex(ID)
	rewardId, _ := params.Args["rewardID"].(string)
	rwId, _ := primitive.ObjectIDFromHex(rewardId)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pkgID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return false, err
	}

	var reward coredb.Reward
	err = db.GetRewardsCollection().FindOne(ctx, map[string]primitive.ObjectID{
		"_id": rwId,
	}).Decode(&reward)
	if err != nil {
		log.Printf("failed to find reward: %v\n", err)
		return nil, err
	}

	var gameSession coredb.GameSession
	err = db.GetRewardsCollection().FindOne(ctx, map[string]primitive.ObjectID{
		"_id": gameId,
	}).Decode(&gameSession)
	if err != nil {
		return false, err
	}

	filter := bson.M{"_id": pkgID}
	update := bson.M{
		"$addToSet": bson.M{"rewards": rewardId},
	}

	_, err = r.GameSessionsRepo.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return false, err
	}

	return true, nil
}
