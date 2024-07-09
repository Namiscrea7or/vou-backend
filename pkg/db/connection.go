package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	db             *mongo.Database
	UserCollection = "users"
)

func GetDB() *mongo.Database {
	if db != nil {
		return db
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	mongoDatabase := os.Getenv("MONGO_DATABASE_NAME")
	connectionURI := os.Getenv("MONGO_URI")

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionURI))
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("connect to db successfully")
	}

	db = client.Database(mongoDatabase)

	return db
}

func GetUsersCollection() *mongo.Collection {
	return GetDB().Collection(UserCollection)
}
