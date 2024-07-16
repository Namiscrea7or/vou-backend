package brand

import (
	"github.com/graphql-go/graphql"
)

type BranchQuery struct {
	BrandRequest    *graphql.Field
	AllBrandRequest *graphql.Field
}

func InitBrandQuery(r *BrandResolver) *BranchQuery {
	return &BranchQuery{
		BrandRequest: &graphql.Field{
			Type:        brandType,
			Description: "Get branch requests by id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: r.GetBrandByID,
		},
		AllBrandRequest: &graphql.Field{
			Type:        graphql.NewList(brandType),
			Description: "Get all branch",
			Resolve:     r.GetAllBrands,
		},
	}
}
