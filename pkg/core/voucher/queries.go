package voucher

import (
	"github.com/graphql-go/graphql"
)

type VouchersQuery struct {
	Voucher           *graphql.Field
	VoucherByCode     *graphql.Field
	Vouchers          *graphql.Field
	VouchersByBrandId *graphql.Field
	VouchersByUserID  *graphql.Field
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
		VoucherByCode: &graphql.Field{
			Type:        VoucherType,
			Description: "Get a voucher by Code",
			Args: graphql.FieldConfigArgument{
				"code": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: r.GetVoucherByCode,
		},
		Vouchers: &graphql.Field{
			Type:        graphql.NewList(VoucherType),
			Description: "Get all vouchers",
			Resolve:     r.GetAllVouchers,
		},
		VouchersByBrandId: &graphql.Field{
			Type:        graphql.NewList(VoucherType),
			Description: "Get all vouchers by brand ID",
			Args: graphql.FieldConfigArgument{
				"brandId": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: r.GetVouchersByBrandId,
		},
		VouchersByUserID: &graphql.Field{
			Type:        graphql.NewList(VoucherType),
			Description: "Get all vouchers by user ID",
			Args: graphql.FieldConfigArgument{
				"userId": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: r.GetVouchersByUserID,
		},
	}
}
