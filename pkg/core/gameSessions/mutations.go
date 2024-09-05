package gameSessions

import (
	"github.com/graphql-go/graphql"
)

type GameSessionsMutation struct {
	CreateGameSession      *graphql.Field
	AddRewardToGameSession *graphql.Field
	EditGameSession        *graphql.Field
	DeleteGameSession      *graphql.Field
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
			Resolve: r.AddRewardToGameSession,
		},
		EditGameSession: &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Edit an existing game session",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"status": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Boolean),
				},
			},
			Resolve: r.EditGameSession,
		},
		DeleteGameSession: &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Delete a game session",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: r.DeleteGameSession,
		},
	}
}
