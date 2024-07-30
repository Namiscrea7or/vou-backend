package gameSessions

import "github.com/graphql-go/graphql"

var gameSessionType = graphql.NewObject(graphql.ObjectConfig{
	Name: "GameSession",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.ID,
		},
		"startTime": &graphql.Field{
			Type: graphql.DateTime,
		},
		"endTime": &graphql.Field{
			Type: graphql.DateTime,
		},
		"rewards": &graphql.Field{
			Type: graphql.NewList(graphql.String),
		},
		"status": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Boolean),
		},
	},
})
