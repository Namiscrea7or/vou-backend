package users

import (
	"github.com/graphql-go/graphql"
)

type UsersQuery struct {
	User        *graphql.Field
	GetAllUsers *graphql.Field
}

func InitUserQuery(r *UsersResolver) *UsersQuery {
	return &UsersQuery{
		User: &graphql.Field{
			Type:        userType,
			Description: "Get a user by email",
			Args: graphql.FieldConfigArgument{
				"email": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: r.GetUserByEmail,
		},
		GetAllUsers: &graphql.Field{
			Type:        graphql.NewList(userType),
			Description: "Get all user",
			Resolve:     r.GetAllUsers,
		},
	}
}
