package coredb

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Name        string             `json:"name" bson:"name"`
	Email       string             `json:"email" bson:"email"`
	FirebaseUID string             `json:"firebaseUID" bson:"firebase_uid"`
	ImageURL    string             `json:"imageURL" bson:"image_url"`
	CreatedAt   time.Time          `json:"createdAt" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updatedAt" bson:"updated_at"`
}

type Voucher struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Code        string             `json:"code" bson:"code"`
	ImageURL    string             `json:"imageURL" bson:"image_url"`
	Value       float64            `json:"value" bson:"value"`
	Description string             `json:"description" bson:"description"`
	ExpiredDate time.Time          `json:"expired_date" bson:"expired_date"`
}

type Package struct {
	ID            primitive.ObjectID `json:"id" bson:"_id"`
	UserID        string             `json:"user_id" bson:"user_id"`
	Vouchers      []string           `json:"vouchers" bson:"vouchers"`
	AllowExchange bool               `json:"allow_exchange" bson:"allow_exchange"`
}
