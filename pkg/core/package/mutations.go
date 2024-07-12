package packages

import (
	"github.com/graphql-go/graphql"
)

type PackagesMutation struct {
	CreatePackage            *graphql.Field
	AddVoucherToPackage      *graphql.Field
	RemoveVoucherFromPackage *graphql.Field
}

func InitPackageMutation(r *PackagesResolver) *PackagesMutation {
	return &PackagesMutation{
		CreatePackage: &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Create a new package",
			Args: graphql.FieldConfigArgument{
				"user_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"allow_exchange": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Boolean),
				},
			},
			Resolve: r.CreatePackage,
		},
		AddVoucherToPackage: &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Add a voucher to a package",
			Args: graphql.FieldConfigArgument{
				"packageID": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"voucherID": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: r.AddVoucherToPackage,
		},
		RemoveVoucherFromPackage: &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Remove a voucher from a package",
			Args: graphql.FieldConfigArgument{
				"packageID": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"voucherID": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: r.RemoveVoucherFromPackage,
		},
	}
}
