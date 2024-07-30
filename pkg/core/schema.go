// core/schema.go
package core

import (
	"vou/pkg/core/brand"
	"vou/pkg/core/exchange"
	"vou/pkg/core/gameSessions"
	packages "vou/pkg/core/package"
	"vou/pkg/core/rewards"
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

		brandResolver = brand.NewBrandResolver()
		brandQuery    = brand.InitBrandQuery(brandResolver)
		brandMutation = brand.InitBrandsMutation(brandResolver)

		gameSessionResolver = gameSessions.NewGameSessionsResolver()
		gameSessionQuery    = gameSessions.InitGameSessionsQuery(gameSessionResolver)
		gameSessionMutation = gameSessions.InitGameSessionsMutation(gameSessionResolver)

		rewardResolver = rewards.NewRewardsResolver()
		rewardQuery    = rewards.InitRewardsQuery(rewardResolver)
		rewardMutation = rewards.InitRewardsMutation(rewardResolver)
	)

	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"user":               usersQuery.User,
			"voucherById":        vouchersQuery.Voucher,
			"voucherByCode":      vouchersQuery.VoucherByCode,
			"package":            packagesQuery.Package,
			"brandById":          brandQuery.BrandRequest,
			"getAllBrand":        brandQuery.AllBrandRequest,
			"getGameSessionByID": gameSessionQuery.GameSession,
			"getRewardByID":      rewardQuery.Reward,
		},
	})

	rootMutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"registerAccount":                usersMutation.RegisterAccount,
			"createVoucher":                  vouchersMutation.CreateVoucher,
			"createPackage":                  packagesMutation.CreatePackage,
			"addVoucherToPackageById":        packagesMutation.AddRewardToPackageById,
			"removeVoucherFromPackageById":   packagesMutation.RemoveRewardFromPackageById,
			"addVoucherToPackageByCode":      packagesMutation.AddVoucherToPackageByCode,
			"removeVoucherFromPackageByCode": packagesMutation.RemoveVoucherFromPackageByCode,
			"createExchangeRequest":          exchangesMutation.CreateExchangeRequest,
			"addVoucherToExchange":           exchangesMutation.AddRewardToExchange,
			"finalizeExchange":               exchangesMutation.FinalizeExchange,
			"createBrand":                    brandMutation.CreateBrand,
			"createGameSession":              gameSessionMutation.CreateGameSession,
			"AddRewardToGameSession":         gameSessionMutation.AddRewardToGameSession,

			"createReward": rewardMutation.CreateReward,
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
