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
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"description": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"type": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"value": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: r.CreateReward,
		},
	}
}
