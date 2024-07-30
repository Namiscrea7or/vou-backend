package rewards

import (
	"github.com/graphql-go/graphql"
)

type RewardsQuery struct {
	Reward *graphql.Field
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
	}
}
