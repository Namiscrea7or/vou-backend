package voucher

import (
	"github.com/graphql-go/graphql"
)

type VouchersMutation struct {
	CreateVoucher *graphql.Field
	EditVoucher   *graphql.Field
	DeleteVoucher *graphql.Field
}

func InitVoucherMutation(r *VouchersResolver) *VouchersMutation {
	return &VouchersMutation{
		CreateVoucher: &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Create a new voucher",
			Args: graphql.FieldConfigArgument{
				"code": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"brandId": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"imageURL": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"value": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"description": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"expiredDate": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.DateTime),
				},
			},
			Resolve: r.CreateVoucher,
		},
		EditVoucher: &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Edit a voucher",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"brandId": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"code": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"imageURL": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"value": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"description": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"expiredDate": &graphql.ArgumentConfig{
					Type: graphql.DateTime,
				},
			},
			Resolve: r.EditVoucher,
		},

		DeleteVoucher: &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Delete a voucher",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"brandId": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: r.DeleteVoucher,
		},
	}
}
