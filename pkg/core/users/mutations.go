package users

import "github.com/graphql-go/graphql"

type UsersMutation struct {
	RegisterAccount *graphql.Field
	Login           *graphql.Field
}

func InitUserMutation(r *UsersResolver) *UsersMutation {
	return &UsersMutation{
		RegisterAccount: &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Register a new account",
			Resolve:     r.RegisterAccount,
			Args: graphql.FieldConfigArgument{
				"username": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"password": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"phoneNumber": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"role": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"profilePicture": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"dob": &graphql.ArgumentConfig{
					Type: (graphql.DateTime),
				},
				"gender": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Boolean),
				},
				"facebookAccount": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
		},
		Login: &graphql.Field{
			Type: graphql.NewObject(graphql.ObjectConfig{
				Name: "LoginResponse",
				Fields: graphql.Fields{
					"token": &graphql.Field{
						Type: graphql.String,
					},
					"user": &graphql.Field{
						Type: userType,
					},
				},
			}),
			Description: "Login a user and return a JWT token",
			Args: graphql.FieldConfigArgument{
				"username": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"password": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: r.Login,
		},
	}
}
