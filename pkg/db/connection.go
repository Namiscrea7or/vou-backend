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

var db *mongo.Database

func GetDB() *mongo.Database {
	if db != nil {
		return db
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	mongoAddress := os.Getenv("MONGO_ADDRESS")
	mongoDatabase := os.Getenv("MONGO_DATABASE_NAME")
	mongoUser := os.Getenv("MONGO_USER")
	mongoPassword := os.Getenv("MONGO_PASSWORD")
	connectionURI := fmt.Sprintf(
		"mongodb://%s:%s@%s:27017/%s",
		mongoUser,
		mongoPassword,
		mongoAddress,
		mongoDatabase,
	)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionURI))
	if err != nil {
		log.Fatal(err)
	}

	db = client.Database(mongoDatabase)

	return db
}
