package exchange

import (
	"github.com/graphql-go/graphql"
)

type ExchangesMutation struct {
	CreateExchange *graphql.Field
	UpdateExchange *graphql.Field
	AskForExchange *graphql.Field
	DeleteExchange *graphql.Field
}

func InitExchangesMutation(r *ExchangesResolver) *ExchangesMutation {
	return &ExchangesMutation{
		CreateExchange: &graphql.Field{
			Type:        exchangeType,
			Description: "Create a new exchange",
			Args: graphql.FieldConfigArgument{
				"rewardIds":     &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.NewList(graphql.String))},
				"voucherId":     &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"gameSessionId": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			},
			Resolve: r.CreateExchange,
		},
		UpdateExchange: &graphql.Field{
			Type:        exchangeType,
			Description: "Update an existing exchange",
			Args: graphql.FieldConfigArgument{
				"id":            &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.ID)},
				"rewardIds":     &graphql.ArgumentConfig{Type: graphql.NewList(graphql.String)},
				"voucherId":     &graphql.ArgumentConfig{Type: graphql.String},
				"completed":     &graphql.ArgumentConfig{Type: graphql.Boolean},
				"gameSessionId": &graphql.ArgumentConfig{Type: graphql.String},
			},
			Resolve: r.UpdateExchange,
		},
		AskForExchange: &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Ask for an exchange by providing user ID, reward IDs, and voucher ID",
			Args: graphql.FieldConfigArgument{
				"userId":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"rewardIds": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.NewList(graphql.String))},
				"voucherId": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			},
			Resolve: r.AskForExchange,
		},
		DeleteExchange: &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Delete an exchange by ID",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: r.DeleteExchange,
		},
	}
}
