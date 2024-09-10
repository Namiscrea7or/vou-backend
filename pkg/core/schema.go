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
		exchangesQuery    = exchange.InitExchangesQuery(exchangesResolver)
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
			"user":                       usersQuery.User,
			"getAllUsers":                usersQuery.GetAllUsers,
			"getUserById":                usersQuery.GetUserByID,
			"voucherById":                vouchersQuery.Voucher,
			"voucherByCode":              vouchersQuery.VoucherByCode,
			"getVouchersByUserId":        vouchersQuery.VouchersByUserID,
			"getAllVouchers":             vouchersQuery.Vouchers,
			"getAllVouchersByBrandId":    vouchersQuery.VouchersByBrandId,
			"getPackageByUserId":         packagesQuery.PackageByUserID,
			"brandById":                  brandQuery.BrandRequest,
			"getAllBrand":                brandQuery.AllBrandRequest,
			"getGameSessionByID":         gameSessionQuery.GameSession,
			"getAllGameSession":          gameSessionQuery.AllGameSessions,
			"getRewardByID":              rewardQuery.Reward,
			"getAllRewards":              rewardQuery.Rewards,
			"getRewardsByUserId":         rewardQuery.GetRewardByUserID,
			"getRewardsBySessionId":      rewardQuery.GetRewardBySessionID,
			"getAllExchanges":            exchangesQuery.GetAllExchanges,
			"getAllExchangesBySessionID": exchangesQuery.GetAllExchangesBySessionID,
		},
	})

	rootMutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"registerAccount":                usersMutation.RegisterAccount,
			"createVoucher":                  vouchersMutation.CreateVoucher,
			"editVoucher":                    vouchersMutation.EditVoucher,
			"deleteVoucher":                  vouchersMutation.DeleteVoucher,
			"createPackage":                  packagesMutation.CreatePackage,
			"addVoucherToPackageById":        packagesMutation.AddRewardToPackageById,
			"removeVoucherFromPackageById":   packagesMutation.RemoveRewardFromPackageById,
			"addVoucherToPackageByCode":      packagesMutation.AddVoucherToPackageByCode,
			"removeVoucherFromPackageByCode": packagesMutation.RemoveVoucherFromPackageByCode,
			"addRewardToPackageById":         packagesMutation.AddRewardToPackageById,
			"removeRewardFromPackageById":    packagesMutation.RemoveRewardFromPackageById,
			"createExchangeRequest":          exchangesMutation.CreateExchange,
			"updateExchangeRequest":          exchangesMutation.UpdateExchange,
			"askForExchange":                 exchangesMutation.AskForExchange,
			"deleteExchangeRequest":          exchangesMutation.DeleteExchange,
			"createBrand":                    brandMutation.CreateBrand,
			"createGameSession":              gameSessionMutation.CreateGameSession,
			"editGameSession":                gameSessionMutation.EditGameSession,
			"deleteGameSession":              gameSessionMutation.DeleteGameSession,
			"addRewardToGameSession":         gameSessionMutation.AddRewardToGameSession,
			"createReward":                   rewardMutation.CreateReward,
			"editReward":                     rewardMutation.EditReward,
			"deleteReward":                   rewardMutation.DeleteReward,
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
