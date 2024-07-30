package gameSessions

import (
	"context"
	"log"
	"time"

	"vou/pkg/db"
	"vou/pkg/db/coredb"

	"github.com/graphql-go/graphql"
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
