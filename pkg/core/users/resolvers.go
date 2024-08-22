package users

import (
	"context"
	"fmt"
	"log"
	"time"

	"vou/pkg/auth"
	"vou/pkg/db"
	"vou/pkg/db/coredb"

	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type UsersResolver struct {
	UsersRepo *coredb.UsersRepo
}

func NewUsersResolver() *UsersResolver {
	return &UsersResolver{
		UsersRepo: coredb.NewUsersRepo(),
	}
}

func (r *UsersResolver) RegisterAccount(params graphql.ResolveParams) (interface{}, error) {
	authProfile, err := auth.GetProfileByContext(params.Context)
	if err != nil {
		return false, err
	}

	username, ok := params.Args["username"].(string)
	if !ok {
		fmt.Errorf("Don't find username")
	}

	password, ok := params.Args["password"].(string)
	if !ok {
		fmt.Errorf("Don't find password")
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), 8)
	password = string(hashed)

	role, ok := params.Args["role"].(string)
	if !ok {
		fmt.Errorf("Don't find role")
	}

	profilePicture, ok := params.Args["profilePicture"].(string)
	if !ok {
		profilePicture = ""
	}

	dob, ok := params.Args["dob"].(time.Time)
	if !ok {
		fmt.Errorf("Don't find dob")
	}

	phoneNumber, ok := params.Args["phoneNumber"].(string)
	if !ok {
		fmt.Errorf("Don't find phoneNumber")
	}

	gender, ok := params.Args["gender"].(bool)
	if !ok {
		fmt.Errorf("Don't find gender")
	}

	facebookAccount, ok := params.Args["facebookAccount"].(string)
	if !ok {
		fmt.Errorf("Don't find facebook Account")
	}

	user := coredb.User{
		ID:              primitive.NewObjectID(),
		Name:            authProfile.Name,
		Username:        username,
		Password:        password,
		Email:           authProfile.Email,
		PhoneNumber:     phoneNumber,
		Role:            role,
		Status:          true,
		ImageURL:        profilePicture,
		DateOfBirth:     dob,
		Gender:          gender,
		FacebookAccount: facebookAccount,
		FirebaseUID:     authProfile.UID,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = db.GetUsersCollection().InsertOne(ctx, user)
	if err != nil {
		log.Printf("failed to insert user: %v\n", err)
		return false, err
	}

	return true, nil
}

func (r *UsersResolver) GetUserByEmail(params graphql.ResolveParams) (interface{}, error) {
	email, ok := params.Args["email"].(string)
	if !ok {
		return nil, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user coredb.User
	err := db.GetUsersCollection().FindOne(ctx, map[string]string{
		"email": email,
	}).Decode(&user)
	if err != nil {
		log.Printf("failed to find user: %v\n", err)
		return nil, err
	}

	return user, nil
}

func (r *UsersResolver) GetAllUsers(params graphql.ResolveParams) (interface{}, error) {
	user, ok := params.Context.Value(auth.UserKey).(coredb.User)
	if !ok {
		return nil, fmt.Errorf("user not found")
	}

	if user.Role != "admin" {
		return nil, fmt.Errorf("Permission denied")
	}

	users, err := r.UsersRepo.GetAllUsers()
	if err != nil {
		return nil, fmt.Errorf("failed to find users: %v", err)
	}

	return users, nil
}
