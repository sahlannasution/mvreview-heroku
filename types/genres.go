package types

import "github.com/graphql-go/graphql"

var GenreType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Genres",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.ID,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
