package gameSessions

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

type GameSessionsResolver struct {
	GameSessionsRepo *coredb.GameSessionsRepo
}

func NewGameSessionsResolver() *GameSessionsResolver {
	return &GameSessionsResolver{
		GameSessionsRepo: coredb.NewGameSessionsRepo(),
	}
}

func (r *GameSessionsResolver) CreateGameSession(params graphql.ResolveParams) (interface{}, error) {
	user, ok := params.Context.Value(auth.UserKey).(coredb.User)
	if !ok {
		return nil, fmt.Errorf("user not found")
	}

	if user.Role != "admin" {
		return nil, fmt.Errorf("Permission denied")
	}

	name, _ := params.Args["name"].(string)
	brandId, _ := params.Args["brandId"].(string)
	startTime, _ := params.Args["startTime"].(time.Time)
	endTime, _ := params.Args["endTime"].(time.Time)
	img, _ := params.Args["image"].(string)

	gameSession := coredb.GameSession{
		ID:        primitive.NewObjectID(),
		Name:      name,
		BrandId:   brandId,
		ImageURL:  img,
		StartTime: startTime,
		EndTime:   endTime,
		Rewards:   []string{},
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

func (r *GameSessionsResolver) GetAllGameSessions(params graphql.ResolveParams) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := db.GetGameSessionsCollection().Find(ctx, bson.M{})
	if err != nil {
		log.Printf("failed to get game sessions: %v\n", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var gameSessions []coredb.GameSession
	if err = cursor.All(ctx, &gameSessions); err != nil {
		log.Printf("failed to decode game sessions: %v\n", err)
		return nil, err
	}

	return gameSessions, nil
}

func (r *GameSessionsResolver) EditGameSession(params graphql.ResolveParams) (interface{}, error) {
	user, ok := params.Context.Value(auth.UserKey).(coredb.User)
	if !ok {
		return nil, fmt.Errorf("user not found")
	}

	if user.Role != "admin" {
		return nil, fmt.Errorf("Permission denied")
	}

	id, err := primitive.ObjectIDFromHex(params.Args["id"].(string))
	if err != nil {
		return false, err
	}

	status, _ := params.Args["status"].(bool)
	name, _ := params.Args["name"].(string)
	startTime, _ := params.Args["startTime"].(time.Time)
	endTime, _ := params.Args["endTime"].(time.Time)
	img, _ := params.Args["image"].(string)

	updateFields := bson.M{}
	if name != "" {
		updateFields["name"] = name
	}

	updateFields["status"] = status
	if !startTime.IsZero() {
		updateFields["startTime"] = startTime
	}

	if !endTime.IsZero() {
		updateFields["endTime"] = endTime
	}

	if img != "" {
		updateFields["imageURL"] = img
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{
		"$set": updateFields,
	}

	_, err = db.GetGameSessionsCollection().UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		log.Printf("failed to edit game session: %v\n", err)
		return false, err
	}

	return true, nil
}

func (r *GameSessionsResolver) DeleteGameSession(params graphql.ResolveParams) (interface{}, error) {
	id, err := primitive.ObjectIDFromHex(params.Args["id"].(string))
	if err != nil {
		return false, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = db.GetGameSessionsCollection().DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		log.Printf("failed to delete game session: %v\n", err)
		return false, err
	}

	return true, nil
}

func (r *GameSessionsResolver) GetGameSessionByBrandID(params graphql.ResolveParams) (interface{}, error) {
	brandId, ok := params.Args["brandId"].(string)
	if !ok {
		return nil, fmt.Errorf("brandId not provided")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := db.GetGameSessionsCollection().Find(ctx, bson.M{"brandId": brandId})
	if err != nil {
		log.Printf("failed to find game sessions for brandId %s: %v\n", brandId, err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var gameSessions []coredb.GameSession
	if err = cursor.All(ctx, &gameSessions); err != nil {
		log.Printf("failed to decode game sessions: %v\n", err)
		return nil, err
	}

	return gameSessions, nil
}
