package packages

import (
	"github.com/graphql-go/graphql"
)

type PackagesQuery struct {
	PackageByID     *graphql.Field
	PackageByUserID *graphql.Field
}

func InitPackageQuery(r *PackagesResolver) *PackagesQuery {
	return &PackagesQuery{
		PackageByID: &graphql.Field{
			Type:        packageType,
			Description: "Get a package by ID",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: r.GetPackageByID,
		},
		PackageByUserID: &graphql.Field{
			Type:        packageType,
			Description: "Get a package by user ID",
			Args: graphql.FieldConfigArgument{
				"userId": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: r.GetPackageByUserID,
		},
	}
}
