package rewards

import (
	"github.com/graphql-go/graphql"
)

type RewardsMutation struct {
	CreateReward *graphql.Field
}

func InitRewardsMutation(r *RewardsResolver) *RewardsMutation {
	return &RewardsMutation{
		CreateReward: &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Create a new reward",
			Resolve:     r.CreateReward,
		},
	}
}
