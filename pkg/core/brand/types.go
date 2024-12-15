package brand

import (
	"github.com/graphql-go/graphql"
)

var gpsType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Gps",
	Fields: graphql.Fields{
		"latitude":  &graphql.Field{Type: graphql.Float},
		"longitude": &graphql.Field{Type: graphql.Float},
	},
})

var brandType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Brand",
	Fields: graphql.Fields{
		"id":        &graphql.Field{Type: graphql.ID},
		"name":      &graphql.Field{Type: graphql.String},
		"industry":  &graphql.Field{Type: graphql.String},
		"address":   &graphql.Field{Type: graphql.String},
		"location":  &graphql.Field{Type: gpsType},
		"status":    &graphql.Field{Type: graphql.Boolean},
		"creatorId": &graphql.Field{Type: graphql.String},
	},
})
