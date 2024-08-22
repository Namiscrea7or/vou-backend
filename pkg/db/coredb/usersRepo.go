package coredb

import (
	"context"
	"fmt"
	"log"
	"time"

	"vou/pkg/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UsersRepo struct {
	*mongo.Collection
}

func NewUsersRepo() *UsersRepo {
	usersCollection := db.GetUsersCollection()
	_, err := usersCollection.Indexes().CreateMany(context.TODO(), []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "email", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys:    bson.D{{Key: "firebase_uid", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	})
	if err != nil {
		log.Println("failed to create indexes for users collection")
		return nil
	}

	return &UsersRepo{usersCollection}
}

func (r *UsersRepo) GetUserByFirebaseUID(firebaseUID string) (User, error) {
	var user User
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.FindOne(ctx, bson.M{"firebase_uid": firebaseUID}).Decode(&user)

	return user, err
}

func (r *UsersRepo) CreateNewUser(user User) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return r.InsertOne(ctx, user)
}

func (r *UsersRepo) GetUserByEmail(email string) (User, error) {
	var user User
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.FindOne(ctx, bson.M{"email": email}).Decode(&user)

	return user, err
}

func (r *UsersRepo) GetAllUsers() ([]User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var users []User
	cursor, err := r.Find(ctx, primitive.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to find users: %v", err)
	}

	defer cursor.Close(ctx)

	err = cursor.All(ctx, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}
