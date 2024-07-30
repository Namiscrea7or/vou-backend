package exchange

import (
	"context"
	"fmt"
	"log"
	"time"

	"vou/pkg/auth"
	"vou/pkg/db"
	"vou/pkg/db/coredb"

	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (r *ExchangesResolver) CreateExchangeRequest(params graphql.ResolveParams) (interface{}, error) {
	user, ok := params.Context.Value(auth.UserKey).(coredb.User)
	if !ok {
		return nil, fmt.Errorf("user not found")
	}

	if user.Role != "user" {
		return nil, fmt.Errorf("Permission denied")
	}

	firstUserID := params.Args["firstUserId"].(string)
	firstRewardId := params.Args["firstRewardId"].(string)

	exchange := coredb.Exchange{
		ID:           primitive.NewObjectID(),
		FirstUserID:  firstUserID,
		FirstRwardID: firstRewardId,
		CreatedAt:    time.Now(),
		Completed:    false,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := db.GetExchangeCollection().InsertOne(ctx, exchange)
	if err != nil {
		log.Printf("failed to create exchange request: %v\n", err)
		return false, err
	}

	return true, nil
}

func (r *ExchangesResolver) AddVoucherToExchange(params graphql.ResolveParams) (interface{}, error) {
	user, ok := params.Context.Value(auth.UserKey).(coredb.User)
	if !ok {
		return nil, fmt.Errorf("user not found")
	}

	if user.Role != "user" {
		return nil, fmt.Errorf("Permission denied")
	}

	exchangeID := params.Args["exchangeId"].(string)
	secondUserID := params.Args["secondUserId"].(string)
	secondRwardId := params.Args["secondRwardId"].(string)

	exchID, err := primitive.ObjectIDFromHex(exchangeID)
	if err != nil {
		return false, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": exchID, "completed": false}
	update := bson.M{
		"$set": bson.M{"secondUserId": secondUserID, "secondRewardId": secondRwardId},
	}

	_, err = db.GetExchangeCollection().UpdateOne(ctx, filter, update)
	if err != nil {
		log.Printf("failed to add voucher to exchange: %v\n", err)
		return false, err
	}

	return true, nil
}

func (r *ExchangesResolver) FinalizeExchange(params graphql.ResolveParams) (interface{}, error) {
	exchangeID := params.Args["exchangeId"].(string)

	exchID, err := primitive.ObjectIDFromHex(exchangeID)
	if err != nil {
		return false, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var exchange coredb.Exchange
	err = db.GetExchangeCollection().FindOne(ctx, bson.M{"_id": exchID}).Decode(&exchange)
	if err != nil {
		log.Printf("failed to find exchange: %v\n", err)
		return false, err
	}

	if err := r.swapRewards(exchange); err != nil {
		return false, err
	}

	filter := bson.M{"_id": exchID}
	update := bson.M{"$set": bson.M{"completed": true}}
	_, err = db.GetExchangeCollection().UpdateOne(ctx, filter, update)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *ExchangesResolver) GetExchangeRequests(params graphql.ResolveParams) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := db.GetExchangeCollection().Find(ctx, bson.M{"completed": false})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var exchanges []coredb.Exchange
	for cursor.Next(ctx) {
		var exchange coredb.Exchange
		if err := cursor.Decode(&exchange); err != nil {
			return nil, err
		}
		exchanges = append(exchanges, exchange)
	}

	return exchanges, nil
}

func (r *ExchangesResolver) swapRewards(exchange coredb.Exchange) error {
	firstUserId, _ := primitive.ObjectIDFromHex(exchange.FirstUserID)
	secondUserId, _ := primitive.ObjectIDFromHex(exchange.SecondUserID)

	if err := r.PackagesRepo.RemoveVoucherFromPackageByCode(firstUserId, exchange.FirstRwardID); err != nil {
		return err
	}
	if err := r.PackagesRepo.AddVoucherToPackageByCode(firstUserId, exchange.SecondRewardID); err != nil {
		return err
	}

	if err := r.PackagesRepo.RemoveVoucherFromPackageByCode(secondUserId, exchange.SecondRewardID); err != nil {
		return err
	}
	if err := r.PackagesRepo.AddVoucherToPackageByCode(secondUserId, exchange.FirstRwardID); err != nil {
		return err
	}

	return nil
}
