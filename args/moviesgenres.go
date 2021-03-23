package args

import "github.com/graphql-go/graphql"

var AddMoviesGenres = graphql.FieldConfigArgument{
	"movies_id": &graphql.ArgumentConfig{
		Type: graphql.Int,
	},
	"genres_id": &graphql.ArgumentConfig{
		Type: graphql.Int,
	},
}
