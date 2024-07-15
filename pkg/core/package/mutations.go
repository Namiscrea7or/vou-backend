package packages

import (
	"github.com/graphql-go/graphql"
)

type PackagesMutation struct {
	CreatePackage                  *graphql.Field
	AddVoucherToPackageById        *graphql.Field
	RemoveVoucherFromPackageById   *graphql.Field
	AddVoucherToPackageByCode      *graphql.Field
	RemoveVoucherFromPackageByCode *graphql.Field
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
		AddVoucherToPackageById: &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Add a voucher to a package by id",
			Args: graphql.FieldConfigArgument{
				"packageID": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"voucherID": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: r.AddVoucherToPackageById,
		},
		RemoveVoucherFromPackageById: &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Remove a voucher from a package by id",
			Args: graphql.FieldConfigArgument{
				"packageID": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"voucherID": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: r.RemoveVoucherFromPackageById,
		},
		AddVoucherToPackageByCode: &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Add a voucher to a package by code",
			Args: graphql.FieldConfigArgument{
				"packageID": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"voucherCode": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: r.AddVoucherToPackageByCode,
		},
		RemoveVoucherFromPackageByCode: &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Remove a voucher from a package by code",
			Args: graphql.FieldConfigArgument{
				"packageID": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"voucherCode": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: r.RemoveVoucherFromPackageByCode,
		},
	}
}
