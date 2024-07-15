// core/schema.go
package core

import (
	"vou/pkg/core/exchange"
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

		exchangesResolver = exchange.NewExchangesResolver()
		exchangesMutation = exchange.InitExchangesMutation(exchangesResolver)
	)

	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"user":          usersQuery.User,
			"voucherById":   vouchersQuery.Voucher,
			"voucherByCode": vouchersQuery.VoucherByCode,
			"package":       packagesQuery.Package,
		},
	})

	rootMutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"registerAccount":                usersMutation.RegisterAccount,
			"createVoucher":                  vouchersMutation.CreateVoucher,
			"createPackage":                  packagesMutation.CreatePackage,
			"addVoucherToPackageById":        packagesMutation.AddVoucherToPackageById,
			"removeVoucherFromPackageById":   packagesMutation.RemoveVoucherFromPackageById,
			"addVoucherToPackageByCode":      packagesMutation.AddVoucherToPackageByCode,
			"removeVoucherFromPackageByCode": packagesMutation.RemoveVoucherFromPackageByCode,
			"createExchangeRequest":          exchangesMutation.CreateExchangeRequest,
			"addVoucherToExchange":           exchangesMutation.AddVoucherToExchange,
			"finalizeExchange":               exchangesMutation.FinalizeExchange,
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
