package gameSessions

import (
	"github.com/graphql-go/graphql"
)

type GameSessionsQuery struct {
	GameSession          *graphql.Field
	AllGameSessions      *graphql.Field
	GameSessionByBrandId *graphql.Field
}

func InitGameSessionsQuery(r *GameSessionsResolver) *GameSessionsQuery {
	return &GameSessionsQuery{
		GameSession: &graphql.Field{
			Type:        gameSessionType,
			Description: "Get a game session by ID",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: r.GetGameSessionByID,
		},
		AllGameSessions: &graphql.Field{
			Type:        graphql.NewList(gameSessionType),
			Description: "Get all game sessions",
			Resolve:     r.GetAllGameSessions,
		},
		GameSessionByBrandId: &graphql.Field{
			Type:        graphql.NewList(gameSessionType),
			Description: "Get game sessions by Brand ID",
			Args: graphql.FieldConfigArgument{
				"brandId": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: r.GetGameSessionByBrandID,
		},
	}
}
