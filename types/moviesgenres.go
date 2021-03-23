package types

import "github.com/graphql-go/graphql"

var MoviesGenresType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "MoviesGenres",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.ID,
			},
			"movies_id": &graphql.Field{
				Type: graphql.Int,
			},
			"genres_id": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)
