package packages

import (
	"github.com/graphql-go/graphql"
)

var packageType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Package",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.ID,
		},
		"user_id": &graphql.Field{
			Type: graphql.String,
		},
		"vouchers": &graphql.Field{
			Type: graphql.NewList(graphql.String),
		},
		"allow_exchange": &graphql.Field{
			Type: graphql.Boolean,
		},
	},
})
