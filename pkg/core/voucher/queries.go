package voucher

import (
	"github.com/graphql-go/graphql"
)

type VouchersQuery struct {
	Voucher *graphql.Field
}

func InitVoucherQuery(r *VouchersResolver) *VouchersQuery {
	return &VouchersQuery{
		Voucher: &graphql.Field{
			Type:        VoucherType,
			Description: "Get a voucher by ID",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: r.GetVoucherByID,
		},
	}
}
