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
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"brandId": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"image": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"startTime": &graphql.ArgumentConfig{
					Type: graphql.DateTime,
				},
				"endTime": &graphql.ArgumentConfig{
					Type: graphql.DateTime,
				},
				"rewards": &graphql.ArgumentConfig{
					Type: graphql.NewList(graphql.String),
				},
			},
			Resolve: r.CreateGameSession,
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
				"name": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"image": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"status": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Boolean),
				},
				"startTime": &graphql.ArgumentConfig{
					Type: graphql.DateTime,
				},
				"endTime": &graphql.ArgumentConfig{
					Type: graphql.DateTime,
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
