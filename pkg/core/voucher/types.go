package voucher

import "github.com/graphql-go/graphql"

var VoucherType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Voucher",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.ID,
		},
		"code": &graphql.Field{
			Type: graphql.String,
		},
		"imageURL": &graphql.Field{
			Type: graphql.String,
		},
		"value": &graphql.Field{
			Type: graphql.Float,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"expiredDate": &graphql.Field{
			Type: graphql.DateTime,
		},
	},
})
