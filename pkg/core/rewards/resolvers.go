package rewards

import (
	"context"
	"log"
	"time"

	"vou/pkg/db"
	"vou/pkg/db/coredb"

	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RewardsResolver struct {
	RewardsRepo *coredb.RewardsRepo
}

func NewRewardsResolver() *RewardsResolver {
	return &RewardsResolver{
		RewardsRepo: coredb.NewRewardsRepo(),
	}
}

func (r *RewardsResolver) CreateReward(params graphql.ResolveParams) (interface{}, error) {
	reward := coredb.Reward{
		ID:          primitive.NewObjectID(),
		Name:        params.Args["name"].(string),
		Description: params.Args["description"].(string),
		Type:        params.Args["type"].(string),
		Value:       params.Args["value"].(string),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := db.GetRewardsCollection().InsertOne(ctx, reward)
	if err != nil {
		log.Printf("failed to insert reward: %v\n", err)
		return false, err
	}

	return true, nil
}

func (r *RewardsResolver) GetRewardByID(params graphql.ResolveParams) (interface{}, error) {
	id, err := primitive.ObjectIDFromHex(params.Args["id"].(string))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var reward coredb.Reward
	err = db.GetRewardsCollection().FindOne(ctx, map[string]primitive.ObjectID{
		"_id": id,
	}).Decode(&reward)
	if err != nil {
		log.Printf("failed to find reward: %v\n", err)
		return nil, err
	}

	return reward, nil
}
