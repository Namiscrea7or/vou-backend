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
	ExpiredDate time.Time          `json:"expiredDate" bson:"expired_date"`
}

type Reward struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	Type        string             `json:"type" bson:"type"`
	Value       string             `json:"value" bson:"value"`
}

type Package struct {
	ID            primitive.ObjectID `json:"id" bson:"_id"`
	UserID        string             `json:"userId" bson:"user_id"`
	Vouchers      []string           `json:"vouchers" bson:"vouchers"`
	Rewards       []string           `json:"rewards" bson:"rewards"`
	AllowExchange bool               `json:"allowExchange" bson:"allow_exchange"`
}

type Exchange struct {
	ID                primitive.ObjectID `json:"id" bson:"_id"`
	FirstUserID       string             `json:"firstUserId" bson:"first_user_id"`
	FirstVoucherCode  string             `json:"firstVoucherCode" bson:"first_voucher_code"`
	SecondUserID      string             `json:"secondUserId" bson:"second_user_id"`
	SecondVoucherCode string             `json:"secondVoucherCode" bson:"second_voucher_code"`
	CreatedAt         time.Time          `json:"createdAt" bson:"created_at"`
	Completed         bool               `json:"completed" bson:"completed"`
}

type Gps struct {
	Latitude  float64 `json:"latitude" bson:"latitude"`
	Longitude float64 `json:"longitude" bson:"longitude"`
}

type Brand struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Name      string             `json:"name" bson:"name"`
	Industry  string             `json:"industry" bson:"industry"`
	Address   string             `json:"address" bson:"address"`
	Location  Gps                `json:"location" bson:"location"`
	Status    bool               `json:"status" bson:"status"`
	CreatorId string             `json:"creatorId" bson:"creator_id"`
}

type GameSession struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	StartTime time.Time          `json:"startTime" bson:"start_time"`
	EndTime   time.Time          `json:"endTime,omitempty" bson:"end_time,omitempty"`
	Status    bool               `json:"status" bson:"status"`
	Rewards   []string           `json:"rewards" bson:"rewards"`
}

type UserGameState struct {
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID        primitive.ObjectID `json:"userId" bson:"user_id"`
	ClaimedReward []ClaimedReward    `json:"claimedRewards" bson:"claimed_rewards"`
}

type ClaimedReward struct {
	ClaimedDate time.Time `json:"claimedDate" bson:"claimed_date"`
	RewardId    string    `json:"rewardId" bson:"reward_id"`
}
