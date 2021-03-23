package args

import "github.com/graphql-go/graphql"

var AddMovies = graphql.FieldConfigArgument{
	"title": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"year": &graphql.ArgumentConfig{
		Type: graphql.Int,
	},
	"ratings": &graphql.ArgumentConfig{
		Type: graphql.Int,
	},
}
