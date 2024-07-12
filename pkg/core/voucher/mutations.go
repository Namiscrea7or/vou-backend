package voucher

import (
	"github.com/graphql-go/graphql"
)

type VouchersMutation struct {
	CreateVoucher *graphql.Field
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
				"imageURL": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"value": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Float),
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
	}
}
