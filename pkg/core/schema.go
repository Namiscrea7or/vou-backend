package core

import (
	packages "vou/pkg/core/package"
	"vou/pkg/core/users"
	"vou/pkg/core/voucher"

	"github.com/graphql-go/graphql"
)

func InitSchema() graphql.Schema {
	var (
		usersResolver = users.NewUsersResolver()
		usersQuery    = users.InitUserQuery(usersResolver)
		usersMutation = users.InitUserMutation(usersResolver)

		vouchersResolver = voucher.NewVouchersResolver()
		vouchersQuery    = voucher.InitVoucherQuery(vouchersResolver)
		vouchersMutation = voucher.InitVoucherMutation(vouchersResolver)

		packagesResolver = packages.NewPackagesResolver()
		packagesQuery    = packages.InitPackageQuery(packagesResolver)
		packagesMutation = packages.InitPackageMutation(packagesResolver)
	)

	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"user":    usersQuery.User,
			"voucher": vouchersQuery.Voucher,
			"package": packagesQuery.Package,
		},
	})

	rootMutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"registerAccount":          usersMutation.RegisterAccount,
			"createVoucher":            vouchersMutation.CreateVoucher,
			"createPackage":            packagesMutation.CreatePackage,
			"addVoucherToPackage":      packagesMutation.AddVoucherToPackage,
			"removeVoucherFromPackage": packagesMutation.RemoveVoucherFromPackage,
		},
	})

	CoreSchema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	})
	if err != nil {
		panic("failed to create schema, error: " + err.Error())
	}

	return CoreSchema
}
