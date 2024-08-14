package users

import "github.com/graphql-go/graphql"

var userType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.ID,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"username": &graphql.Field{
			Type: graphql.String,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
		"phoneNumber": &graphql.Field{
			Type: graphql.String,
		},
		"role": &graphql.Field{
			Type: graphql.String,
		},
		"status": &graphql.Field{
			Type: graphql.Boolean,
		},
		"imageURL": &graphql.Field{
			Type: graphql.String,
		},
		"dateOfBirth": &graphql.Field{
			Type: graphql.String,
		},
		"gender": &graphql.Field{
			Type: graphql.Boolean,
		},
		"facebookAccount": &graphql.Field{
			Type: graphql.String,
		},
		"firebaseUID": &graphql.Field{
			Type: graphql.String,
		},
		"createdAt": &graphql.Field{
			Type: graphql.DateTime,
		},
		"updatedAt": &graphql.Field{
			Type: graphql.DateTime,
		},
	},
})
