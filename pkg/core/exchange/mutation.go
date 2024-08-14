package exchange

import (
	"github.com/graphql-go/graphql"
)

type ExchangesMutation struct {
	CreateExchangeRequest *graphql.Field
	AddRewardToExchange   *graphql.Field
	FinalizeExchange      *graphql.Field
}

func InitExchangesMutation(r *ExchangesResolver) *ExchangesMutation {
	return &ExchangesMutation{
		CreateExchangeRequest: &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Create a new exchange request",
			Args: graphql.FieldConfigArgument{
				"firstUserId": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"firstRewardId": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: r.CreateExchangeRequest,
		},
		AddRewardToExchange: &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Add a voucher to an existing exchange request",
			Args: graphql.FieldConfigArgument{
				"exchangeId": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"secondUserId": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"secondRewardId": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: r.AddRewardToExchange,
		},
		FinalizeExchange: &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Finalize an exchange request",
			Args: graphql.FieldConfigArgument{
				"exchangeId": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: r.FinalizeExchange,
		},
	}
}
