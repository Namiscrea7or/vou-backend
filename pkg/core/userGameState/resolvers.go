package userGameState

import (
	"context"
	"log"
	"time"

	"vou/pkg/db"
	"vou/pkg/db/coredb"

	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserGameStatesResolver struct {
	UserGameStatesRepo *coredb.UserGameStatesRepo
}

func NewUserGameStatesResolver() *UserGameStatesResolver {
	return &UserGameStatesResolver{
		UserGameStatesRepo: coredb.NewUserGameStatesRepo(),
	}
}

func (r *UserGameStatesResolver) UpdateUserGameState(params graphql.ResolveParams) (interface{}, error) {
	userID, err := primitive.ObjectIDFromHex(params.Args["userId"].(string))
	if err != nil {
		return false, err
	}

	claimedReward := coredb.ClaimedReward{
		RewardId:    "",
		ClaimedDate: time.Now(),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := map[string]primitive.ObjectID{
		"user_id": userID,
	}

	update := map[string]interface{}{
		"$push": map[string]interface{}{
			"claimed_rewards": claimedReward,
		},
	}

	_, err = db.GetUserGameStatesCollection().UpdateOne(ctx, filter, update)
	if err != nil {
		log.Printf("failed to update user game state: %v\n", err)
		return false, err
	}

	return true, nil
}

func (r *UserGameStatesResolver) GetUserGameStateByUserID(params graphql.ResolveParams) (interface{}, error) {
	userID, err := primitive.ObjectIDFromHex(params.Args["userId"].(string))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var userGameState coredb.UserGameState
	err = db.GetUserGameStatesCollection().FindOne(ctx, map[string]primitive.ObjectID{
		"user_id": userID,
	}).Decode(&userGameState)
	if err != nil {
		log.Printf("failed to find user game state: %v\n", err)
		return nil, err
	}

	return userGameState, nil
}
