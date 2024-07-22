package userGameState

import (
	"github.com/graphql-go/graphql"
)

type UserGameStatesQuery struct {
	UserGameState *graphql.Field
}

func InitUserGameStatesQuery(r *UserGameStatesResolver) *UserGameStatesQuery {
	return &UserGameStatesQuery{
		UserGameState: &graphql.Field{
			Type:        userGameStateType,
			Description: "Get user game state by user ID",
			Args: graphql.FieldConfigArgument{
				"userId": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: r.GetUserGameStateByUserID,
		},
	}
}
