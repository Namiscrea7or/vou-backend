package gameSessions

import (
	"github.com/graphql-go/graphql"
)

type GameSessionsMutation struct {
	CreateGameSession *graphql.Field
}

func InitGameSessionsMutation(r *GameSessionsResolver) *GameSessionsMutation {
	return &GameSessionsMutation{
		CreateGameSession: &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Create a new game session",
			Resolve:     r.CreateGameSession,
		},
	}
}
