package rewards

import (
	"github.com/graphql-go/graphql"
)

type RewardsMutation struct {
	CreateReward *graphql.Field
	EditReward   *graphql.Field
	DeleteReward *graphql.Field
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
				"gameId": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"image": &graphql.ArgumentConfig{
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
		EditReward: &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Edit an existing reward",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"image": &graphql.ArgumentConfig{
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
			Resolve: r.EditReward,
		},
		DeleteReward: &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Delete a reward",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: r.DeleteReward,
		},
	}
}
