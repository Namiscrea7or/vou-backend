package packages

import (
	"github.com/graphql-go/graphql"
)

type PackagesQuery struct {
	Package *graphql.Field
}

func InitPackageQuery(r *PackagesResolver) *PackagesQuery {
	return &PackagesQuery{
		Package: &graphql.Field{
			Type:        packageType,
			Description: "Get a package by ID",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: r.GetPackageByID,
		},
	}
}
