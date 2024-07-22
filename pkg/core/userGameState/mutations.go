package userGameState

import (
	"github.com/graphql-go/graphql"
)

type UserGameStatesMutation struct {
	UpdateUserGameState *graphql.Field
}

func InitUserGameStatesMutation(r *UserGameStatesResolver) *UserGameStatesMutation {
	return &UserGameStatesMutation{
		UpdateUserGameState: &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Update user game state",
			Resolve:     r.UpdateUserGameState,
		},
	}
}
