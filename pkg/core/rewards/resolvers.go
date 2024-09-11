package rewards

import (
	"context"
	"fmt"
	"log"
	"time"

	"vou/pkg/db"
	"vou/pkg/db/coredb"

	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
		ImageURL:    params.Args["image"].(string),
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

	gameSessionID, err := primitive.ObjectIDFromHex(params.Args["gameId"].(string))
	if err != nil {
		log.Printf("invalid game session ID: %v\n", err)
		return false, err
	}

	filter := bson.M{"_id": gameSessionID}
	update := bson.M{
		"$addToSet": bson.M{"rewards": reward.ID.Hex()},
	}

	_, err = db.GetGameSessionsCollection().UpdateOne(ctx, filter, update)
	if err != nil {
		log.Printf("failed to add reward to game session: %v\n", err)
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
	err = db.GetRewardsCollection().FindOne(ctx, bson.M{"_id": id}).Decode(&reward)
	if err != nil {
		log.Printf("failed to find reward: %v\n", err)
		return nil, err
	}

	return reward, nil
}

func (r *RewardsResolver) GetAllRewards(params graphql.ResolveParams) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := db.GetRewardsCollection().Find(ctx, bson.M{})
	if err != nil {
		log.Printf("failed to fetch rewards: %v\n", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var rewards []coredb.Reward
	for cursor.Next(ctx) {
		var reward coredb.Reward
		if err = cursor.Decode(&reward); err != nil {
			log.Printf("failed to decode reward: %v\n", err)
			return nil, err
		}
		rewards = append(rewards, reward)
	}

	return rewards, nil
}

func (r *RewardsResolver) EditReward(params graphql.ResolveParams) (interface{}, error) {
	id, err := primitive.ObjectIDFromHex(params.Args["id"].(string))
	if err != nil {
		return nil, err
	}

	update := bson.M{
		"$set": bson.M{
			"name":        params.Args["name"].(string),
			"imageURL":    params.Args["image"].(string),
			"description": params.Args["description"].(string),
			"type":        params.Args["type"].(string),
			"value":       params.Args["value"].(string),
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = db.GetRewardsCollection().UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		log.Printf("failed to update reward: %v\n", err)
		return false, err
	}

	return true, nil
}

func (r *RewardsResolver) DeleteReward(params graphql.ResolveParams) (interface{}, error) {
	id, err := primitive.ObjectIDFromHex(params.Args["id"].(string))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = db.GetRewardsCollection().DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		log.Printf("failed to delete reward: %v\n", err)
		return false, err
	}

	return true, nil
}

func (r *RewardsResolver) GetRewardBySessionID(params graphql.ResolveParams) (interface{}, error) {
	sessionID, err := primitive.ObjectIDFromHex(params.Args["sessionId"].(string))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var gameSession coredb.GameSession
	err = db.GetGameSessionsCollection().FindOne(ctx, bson.M{"_id": sessionID}).Decode(&gameSession)
	if err != nil {
		log.Printf("failed to find game session: %v\n", err)
		return nil, err
	}

	if len(gameSession.Rewards) == 0 {
		return []coredb.Reward{}, nil
	}

	rewardIDs := make([]primitive.ObjectID, len(gameSession.Rewards))
	for i, id := range gameSession.Rewards {
		objectID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			log.Printf("invalid reward ID: %v\n", err)
			return nil, err
		}
		rewardIDs[i] = objectID
	}

	cursor, err := db.GetRewardsCollection().Find(ctx, bson.M{"_id": bson.M{"$in": rewardIDs}})
	if err != nil {
		log.Printf("failed to fetch rewards: %v\n", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var rewards []coredb.Reward
	for cursor.Next(ctx) {
		var reward coredb.Reward
		if err = cursor.Decode(&reward); err != nil {
			log.Printf("failed to decode reward: %v\n", err)
			return nil, err
		}
		rewards = append(rewards, reward)
	}

	if err = cursor.Err(); err != nil {
		log.Printf("cursor error: %v\n", err)
		return nil, err
	}

	return rewards, nil
}

func (r *RewardsResolver) GetRewardByUserID(params graphql.ResolveParams) (interface{}, error) {
	userID, ok := params.Args["userId"].(string)
	if !ok {
		return nil, fmt.Errorf("invalid user ID")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var pkg coredb.Package
	err := db.GetPackageCollection().FindOne(ctx, bson.M{"user_id": userID}).Decode(&pkg)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return []coredb.Reward{}, nil
		}
		log.Printf("failed to fetch package: %v\n", err)
		return nil, err
	}

	log.Printf("pkg.Rewards: %v", pkg.Rewards)

	var rewards []coredb.Reward

	for _, rewardIDStr := range pkg.Rewards {
		objectID, err := primitive.ObjectIDFromHex(rewardIDStr)
		if err != nil {
			log.Printf("invalid reward ID: %v\n", err)
			return nil, err
		}

		var reward coredb.Reward
		err = db.GetRewardsCollection().FindOne(ctx, bson.M{"_id": objectID}).Decode(&reward)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				continue
			}
			log.Printf("failed to fetch reward: %v\n", err)
			return nil, err
		}

		rewards = append(rewards, reward)
	}

	log.Printf("Fetched rewards: %v", rewards)

	return rewards, nil
}
