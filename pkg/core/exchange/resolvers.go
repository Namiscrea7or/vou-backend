package exchange

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

type ExchangesResolver struct {
	ExchangesRepo *coredb.ExchangesRepo
	PackagesRepo  *coredb.PackagesRepo
}

func NewExchangesResolver() *ExchangesResolver {
	return &ExchangesResolver{
		ExchangesRepo: coredb.NewExchangesRepo(),
		PackagesRepo:  coredb.NewPackagesRepo(),
	}
}

func (r *ExchangesResolver) CreateExchange(params graphql.ResolveParams) (interface{}, error) {
	rewardIds := castToStringSlice(params.Args["rewardIds"].([]interface{}))
	voucherId := params.Args["voucherId"].(string)
	gameSessionId := params.Args["gameSessionId"].(string)

	exchange := coredb.Exchange{
		ID:            primitive.NewObjectID(),
		RewardIDs:     rewardIds,
		VoucherID:     voucherId,
		GameSessionID: gameSessionId,
		CreatedAt:     time.Now(),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := db.GetExchangeCollection().InsertOne(ctx, exchange)
	if err != nil {
		log.Printf("failed to create exchange: %v\n", err)
		return nil, err
	}

	return exchange, nil
}

func (r *ExchangesResolver) GetExchangeByID(params graphql.ResolveParams) (interface{}, error) {
	id, err := primitive.ObjectIDFromHex(params.Args["id"].(string))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var exchange coredb.Exchange
	err = db.GetExchangeCollection().FindOne(ctx, bson.M{"_id": id}).Decode(&exchange)
	if err != nil {
		log.Printf("failed to find exchange: %v\n", err)
		return nil, err
	}

	return exchange, nil
}

func (r *ExchangesResolver) GetAllExchanges(params graphql.ResolveParams) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := db.GetExchangeCollection().Find(ctx, bson.M{})
	if err != nil {
		log.Printf("failed to fetch exchanges: %v\n", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var exchanges []coredb.Exchange
	for cursor.Next(ctx) {
		var exchange coredb.Exchange
		if err = cursor.Decode(&exchange); err != nil {
			log.Printf("failed to decode exchange: %v\n", err)
			return nil, err
		}
		exchanges = append(exchanges, exchange)
	}

	if err = cursor.Err(); err != nil {
		log.Printf("cursor error: %v\n", err)
		return nil, err
	}

	return exchanges, nil
}

func (r *ExchangesResolver) UpdateExchange(params graphql.ResolveParams) (interface{}, error) {
	id, err := primitive.ObjectIDFromHex(params.Args["id"].(string))
	if err != nil {
		return nil, err
	}

	update := bson.M{
		"$set": bson.M{
			"reward_ids":     castToStringSlice(params.Args["rewardIds"].([]interface{})),
			"voucher_id":     params.Args["voucherId"].(string),
			"gameSession_id": params.Args["gameSessionId"].(string),
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = db.GetExchangeCollection().UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		log.Printf("failed to update exchange: %v\n", err)
		return false, err
	}

	return true, nil
}

func (r *ExchangesResolver) DeleteExchange(params graphql.ResolveParams) (interface{}, error) {
	id, err := primitive.ObjectIDFromHex(params.Args["id"].(string))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = db.GetExchangeCollection().DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		log.Printf("failed to delete exchange: %v\n", err)
		return false, err
	}

	return true, nil
}

func (r *ExchangesResolver) GetAllExchangesByGameSessionID(params graphql.ResolveParams) (interface{}, error) {
	sessionID, err := primitive.ObjectIDFromHex(params.Args["sessionId"].(string))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := db.GetExchangeCollection().Find(ctx, bson.M{"gameSession_id": sessionID.Hex()})
	if err != nil {
		log.Printf("failed to fetch exchanges: %v\n", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var exchanges []coredb.Exchange
	for cursor.Next(ctx) {
		var exchange coredb.Exchange
		if err = cursor.Decode(&exchange); err != nil {
			log.Printf("failed to decode exchange: %v\n", err)
			return nil, err
		}
		exchanges = append(exchanges, exchange)
	}

	if err = cursor.Err(); err != nil {
		log.Printf("cursor error: %v\n", err)
		return nil, err
	}

	return exchanges, nil
}

func castToStringSlice(i interface{}) []string {
	var result []string
	if slice, ok := i.([]interface{}); ok {
		for _, item := range slice {
			if str, ok := item.(string); ok {
				result = append(result, str)
			}
		}
	}
	return result
}

func (r *ExchangesResolver) AskForExchange(params graphql.ResolveParams) (interface{}, error) {
	userId := params.Args["userId"].(string)
	rewardIds := castToStringSlice(params.Args["rewardIds"].([]interface{}))
	voucherId := params.Args["voucherId"].(string)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var exchange coredb.Exchange
	err := db.GetExchangeCollection().FindOne(ctx, bson.M{
		"reward_ids": bson.M{"$in": rewardIds},
		"voucher_id": voucherId,
	}).Decode(&exchange)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Printf("no exchange found for rewards %v and voucher %s\n", rewardIds, voucherId)
			return nil, fmt.Errorf("no matching exchange found")
		}
		log.Printf("failed to fetch exchange: %v\n", err)
		return nil, err
	}

	var pkg coredb.Package
	err = db.GetPackageCollection().FindOne(ctx, bson.M{"user_id": userId}).Decode(&pkg)
	if err != nil {
		log.Printf("failed to fetch package: %v\n", err)
		return nil, err
	}

	for _, rewardID := range exchange.RewardIDs {
		found := false
		for i, r := range pkg.Rewards {
			if r == rewardID {
				pkg.Rewards = append(pkg.Rewards[:i], pkg.Rewards[i+1:]...)
				found = true
				break
			}
		}

		if !found {
			log.Printf("reward %s not found in user's package\n", rewardID)
			return nil, fmt.Errorf("reward %s not found in user's package", rewardID)
		}
	}

	pkg.Vouchers = append(pkg.Vouchers, exchange.VoucherID)

	_, err = db.GetPackageCollection().UpdateOne(
		ctx,
		bson.M{"user_id": userId},
		bson.M{"$set": bson.M{"rewards": pkg.Rewards, "vouchers": pkg.Vouchers}},
	)
	if err != nil {
		log.Printf("failed to update user's package: %v\n", err)
		return nil, err
	}

	return true, nil
}
