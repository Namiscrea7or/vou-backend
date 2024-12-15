package userGameState

import "github.com/graphql-go/graphql"

var claimedRewardType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ClaimedReward",
	Fields: graphql.Fields{
		"claimedDate": &graphql.Field{
			Type: graphql.NewNonNull(graphql.DateTime),
		},
		"rewardId": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
})

var userGameStateType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserGameState",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.ID,
		},
		"userId": &graphql.Field{
			Type: graphql.ID,
		},
		"claimedRewards": &graphql.Field{
			Type: graphql.NewList(claimedRewardType),
		},
	},
})
