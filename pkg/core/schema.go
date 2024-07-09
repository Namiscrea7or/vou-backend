package core

import (
	"vou/pkg/core/users"

	"github.com/graphql-go/graphql"
)

func InitSchema() graphql.Schema {
	var (
		usersResolver = users.NewUsersResolver()
		usersQuery    = users.InitUserQuery(usersResolver)
		usersMutation = users.InitUserMutation(usersResolver)
	)
	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"user": usersQuery.User,
		},
	})

	rootMutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"registerAccount": usersMutation.RegisterAccount,
		},
	})

	CoreSchema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	})

	return CoreSchema
}
