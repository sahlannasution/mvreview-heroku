package services

import (
	"github.com/sahlannasution/mvreview-heroku/args"
	"github.com/sahlannasution/mvreview-heroku/resolver"
	"github.com/sahlannasution/mvreview-heroku/types"

	"github.com/graphql-go/graphql"
)

var QueryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			// getAllUser Data
			// URL : localhost:8080/?Query={getAllUser{id, email, fullName, role}}
			"getAllUser": &graphql.Field{
				Type:    graphql.NewList(types.UserType), // Define User Type
				Resolve: resolver.GetAllUser,             // Call getAllUser resolver function
			},

			// getUser By Id
			// url localhost:8080/?Query={getUser(email:"sahlan.nasution07@gmail.com"){id, email, fullName, role}}
			"getUser": &graphql.Field{
				Type:    types.UserType,   // Define User Type
				Args:    args.GetUserArgs, // Define Arguments to getUser
				Resolve: resolver.GetUser, // Call getUser resolver function
			},

			// getGenreList
			// url localhost:8080/?Query={getGenreList{id, name}}
			"getGenreList": &graphql.Field{
				Type:    graphql.NewList(types.GenreType), // Define Genre Type
				Resolve: resolver.GetGenresList,           // Call GetGenresList resolver function
			},

			// getMovieList
			// url localhost:8080/?Query={getMovieList{id, title, year, ratings, genres:{id, name}}}
			"getMovieList": &graphql.Field{
				Type:    graphql.NewList(types.MoviesType), // Define Genre Type
				Resolve: resolver.GetMoviesList,            // Call GetGenresList resolver function
			},

			// getReviewDetail
			// url localhost:8080/?Query={getReviewDetail(id:1){id, review, rate, users_id, movies_id}}
			"getReviewDetail": &graphql.Field{
				Type:    types.ReviewType,         // Define Review Type
				Args:    args.GetReviewDetail,     // Define Arguments to getReviewDetail
				Resolve: resolver.GetReviewDetail, // Call GetReviewDetail resolver function
			},

			// getReviewByMovie
			// url localhost:8080/?Query={getReviewByMovie(movies_id:1){id, review, rate, users_id, movies_id}}
			"getReviewByMovie": &graphql.Field{
				Type:    graphql.NewList(types.ReviewType), // Define Review Type
				Args:    args.GetReviewByMovie,             // Define Arguments to getReviewByMovie
				Resolve: resolver.GetReviewByMovie,         // Call GetReviewByMovie resolver function
			},
		},
	},
)
