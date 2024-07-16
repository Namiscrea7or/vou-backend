package exchange

import (
	"github.com/graphql-go/graphql"
)

var exchangeType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Exchange",
	Fields: graphql.Fields{
		"id":                &graphql.Field{Type: graphql.ID},
		"firstUserId":       &graphql.Field{Type: graphql.String},
		"firstVoucherCode":  &graphql.Field{Type: graphql.String},
		"secondUserId":      &graphql.Field{Type: graphql.String},
		"secondVoucherCode": &graphql.Field{Type: graphql.String},
		"createdAt":         &graphql.Field{Type: graphql.DateTime},
		"completed":         &graphql.Field{Type: graphql.Boolean},
	},
})
