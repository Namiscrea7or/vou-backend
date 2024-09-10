package exchange

import (
	"github.com/graphql-go/graphql"
)

var exchangeType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Exchange",
	Fields: graphql.Fields{
		"id":            &graphql.Field{Type: graphql.ID},
		"gameSessionId": &graphql.Field{Type: graphql.String},
		"rewardIds":     &graphql.Field{Type: graphql.NewList(graphql.String)},
		"voucherId":     &graphql.Field{Type: graphql.String},
		"createdAt":     &graphql.Field{Type: graphql.DateTime},
	},
})
