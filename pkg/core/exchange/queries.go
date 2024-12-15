package exchange

import (
	"github.com/graphql-go/graphql"
)

type ExchangesQuery struct {
	ExchangeRequests *graphql.Field
}

func InitExchangesQuery(r *ExchangesResolver) *ExchangesQuery {
	return &ExchangesQuery{
		ExchangeRequests: &graphql.Field{
			Type:        graphql.NewList(exchangeType),
			Description: "Get all exchange requests",
			Resolve:     r.GetExchangeRequests,
		},
	}
}
