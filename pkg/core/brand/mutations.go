package brand

import (
	"github.com/graphql-go/graphql"
)

type BrandMutation struct {
	CreateBrand *graphql.Field
}

func InitBrandsMutation(r *BrandResolver) *BrandMutation {
	return &BrandMutation{
		CreateBrand: &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Create a new Branch",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"industry": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"address": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"latiude": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Float),
				},
				"longiude": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Float),
				},
				"status": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Boolean),
				},
				"creatorId": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: r.CreateBrand,
		},
	}
}
