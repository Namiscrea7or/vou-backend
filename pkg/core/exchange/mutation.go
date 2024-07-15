package exchange

import (
	"github.com/graphql-go/graphql"
)

type ExchangesMutation struct {
	CreateExchangeRequest *graphql.Field
	AddVoucherToExchange  *graphql.Field
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
				"firstVoucherCode": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: r.CreateExchangeRequest,
		},
		AddVoucherToExchange: &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Add a voucher to an existing exchange request",
			Args: graphql.FieldConfigArgument{
				"exchangeId": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"secondUserId": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"secondVoucherCode": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: r.AddVoucherToExchange,
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
