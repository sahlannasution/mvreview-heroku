package args

import "github.com/graphql-go/graphql"

var GetReviewDetail = graphql.FieldConfigArgument{
	"id": &graphql.ArgumentConfig{
		Type: graphql.Int,
	},
}

var GetReviewByMovie = graphql.FieldConfigArgument{
	"movies_id": &graphql.ArgumentConfig{
		Type: graphql.Int,
	},
}

var AddReview = graphql.FieldConfigArgument{
	"review": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"rate": &graphql.ArgumentConfig{
		Type: graphql.Int,
	},
	"movies_id": &graphql.ArgumentConfig{
		Type: graphql.Int,
	},
	"users_id": &graphql.ArgumentConfig{
		Type: graphql.Int,
	},
}
