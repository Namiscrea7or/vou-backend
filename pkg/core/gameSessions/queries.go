package gameSessions

import (
	"github.com/graphql-go/graphql"
)

type GameSessionsQuery struct {
	GameSession *graphql.Field
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
	}
}
