package users

import "github.com/graphql-go/graphql"

type UsersMutation struct {
	RegisterAccount *graphql.Field
}

func InitUserMutation(r *UsersResolver) *UsersMutation {
	return &UsersMutation{
		RegisterAccount: &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Register a new account",
			Resolve:     r.RegisterAccount,
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"username": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"password": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"email": &graphql.ArgumentConfig{
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
	}
}
