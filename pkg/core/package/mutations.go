package packages

import (
	"github.com/graphql-go/graphql"
)

type PackagesMutation struct {
	CreatePackage                  *graphql.Field
	AddRewardToPackageById         *graphql.Field
	RemoveRewardFromPackageById    *graphql.Field
	AddVoucherToPackageByCode      *graphql.Field
	RemoveVoucherFromPackageByCode *graphql.Field
}

func InitPackageMutation(r *PackagesResolver) *PackagesMutation {
	return &PackagesMutation{
		CreatePackage: &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Create a new package",
			Args: graphql.FieldConfigArgument{
				"userId": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"allowExchange": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Boolean),
				},
			},
			Resolve: r.CreatePackage,
		},
		AddRewardToPackageById: &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Add a voucher to a package by id",
			Args: graphql.FieldConfigArgument{
				"packageID": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"rewardID": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: r.AddRewardToPackageById,
		},
		RemoveRewardFromPackageById: &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Remove a voucher from a package by id",
			Args: graphql.FieldConfigArgument{
				"packageID": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"rewardID": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: r.RemoveRewardFromPackageById,
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
