package rewards

import "github.com/graphql-go/graphql"

var rewardType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Reward",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.ID,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"type": &graphql.Field{
			Type: graphql.String,
		},
		"value": &graphql.Field{
			Type: graphql.String,
		},
	},
})
