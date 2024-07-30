package gameSessions

import (
	"github.com/graphql-go/graphql"
)

type GameSessionsMutation struct {
	CreateGameSession      *graphql.Field
	AddRewardToGameSession *graphql.Field
}

func InitGameSessionsMutation(r *GameSessionsResolver) *GameSessionsMutation {
	return &GameSessionsMutation{
		CreateGameSession: &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Create a new game session",
			Resolve:     r.CreateGameSession,
		},
		AddRewardToGameSession: &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Add reward to game session",
			Args: graphql.FieldConfigArgument{
				"gameSessionID": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"rewardID": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
		},
	}
}
