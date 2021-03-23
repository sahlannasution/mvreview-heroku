package args

import "github.com/graphql-go/graphql"

var GetUserArgs = graphql.FieldConfigArgument{
	"email": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
}

var RegisterArgs = graphql.FieldConfigArgument{
	"email": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"password": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"fullName": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"role": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
}
