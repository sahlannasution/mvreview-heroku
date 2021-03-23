package types

import "github.com/graphql-go/graphql"

var ReviewType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Movies",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.ID,
			},
			"review": &graphql.Field{
				Type: graphql.String,
			},
			"rate": &graphql.Field{
				Type: graphql.Int,
			},
			"users_id": &graphql.Field{
				Type: graphql.ID,
			},
			"movies_id": &graphql.Field{
				Type: graphql.ID,
			},
		},
	},
)
