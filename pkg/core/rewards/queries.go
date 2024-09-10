package rewards

import (
	"github.com/graphql-go/graphql"
)

type RewardsQuery struct {
	Reward               *graphql.Field
	Rewards              *graphql.Field
	GetRewardBySessionID *graphql.Field
	GetRewardByUserID    *graphql.Field
}

func InitRewardsQuery(r *RewardsResolver) *RewardsQuery {
	return &RewardsQuery{
		Reward: &graphql.Field{
			Type:        rewardType,
			Description: "Get a reward by ID",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: r.GetRewardByID,
		},
		Rewards: &graphql.Field{
			Type:        graphql.NewList(rewardType),
			Description: "Get all rewards",
			Resolve:     r.GetAllRewards,
		},
		GetRewardBySessionID: &graphql.Field{
			Type:        graphql.NewList(rewardType),
			Description: "Get rewards by session ID",
			Args: graphql.FieldConfigArgument{
				"sessionId": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: r.GetRewardBySessionID,
		},
		GetRewardByUserID: &graphql.Field{
			Type:        graphql.NewList(rewardType),
			Description: "Get rewards by user ID",
			Args: graphql.FieldConfigArgument{
				"userId": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: r.GetRewardByUserID,
		},
	}
}
