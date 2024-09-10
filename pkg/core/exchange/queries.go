package exchange

import (
	"github.com/graphql-go/graphql"
)

type ExchangesQuery struct {
	GetAllExchanges            *graphql.Field
	GetAllExchangesBySessionID *graphql.Field
}

func InitExchangesQuery(r *ExchangesResolver) *ExchangesQuery {
	return &ExchangesQuery{
		GetAllExchanges: &graphql.Field{
			Type:        graphql.NewList(exchangeType),
			Description: "Get all exchanges",
			Resolve:     r.GetAllExchanges,
		},
		GetAllExchangesBySessionID: &graphql.Field{
			Type:        graphql.NewList(exchangeType),
			Description: "Get all exchanges by game session ID",
			Args: graphql.FieldConfigArgument{
				"sessionId": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: r.GetAllExchangesByGameSessionID,
		},
	}
}
