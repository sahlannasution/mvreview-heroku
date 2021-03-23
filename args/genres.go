package args

import "github.com/graphql-go/graphql"

var AddGenre = graphql.FieldConfigArgument{
	"name": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
}
