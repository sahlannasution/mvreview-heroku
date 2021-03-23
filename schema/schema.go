package schema

import (
	"github.com/sahlannasution/mvreview-heroku/services"

	"github.com/graphql-go/graphql"
)

// Define schema for route
var Schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    services.QueryType,    // Define query type for get data
		Mutation: services.MutationType, // Define mutation type for create, update, delete
	},
)
