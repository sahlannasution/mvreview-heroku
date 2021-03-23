package types

import "github.com/graphql-go/graphql"

var MoviesType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Movies",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.ID,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"year": &graphql.Field{
				Type: graphql.Int,
			},
			"ratings": &graphql.Field{
				Type: graphql.Int,
			},
			"genres": &graphql.Field{
				Type: graphql.NewList(GenreType),
			},
		},
	},
)
