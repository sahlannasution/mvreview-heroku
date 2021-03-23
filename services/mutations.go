package services

import (
	"github.com/sahlannasution/mvreview-heroku/args"
	"github.com/sahlannasution/mvreview-heroku/resolver"
	"github.com/sahlannasution/mvreview-heroku/types"

	"github.com/graphql-go/graphql"
)

var MutationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			// URL : localhost:8080/?Query=mutation+_{register(email:"sahlan.nasution07@gmail.com", password:"123456789", fullName:"Sahlan NST", role:"guest"){id, email, fullName, role}}
			"register": &graphql.Field{
				Type:    types.UserType,    // Define UserType
				Args:    args.RegisterArgs, // Define Register Arguments (params to add)
				Resolve: resolver.Register, // Call Register resolver function
			},
			// URL : localhost:8080/?Query=mutation+_{addGenre(name:"Seram"){id, name}}
			"addGenre": &graphql.Field{
				Type:    types.GenreType,   // Define GenreType
				Args:    args.AddGenre,     // Define AddGenre Arguments (params to add)
				Resolve: resolver.AddGenre, // Call AddGenre resolver function
			},
			// URL : localhost:8080/?Query=mutation+_{addMovies(title:"Beranak Dalam Kubur", year:2018, ratings:0){id, title, year, ratings}}
			"addMovies": &graphql.Field{
				Type:    types.MoviesType,   // Define MoviesType
				Args:    args.AddMovies,     // Define AddMovies Arguments (params to add)
				Resolve: resolver.AddMovies, // Call AddMovies resolver function
			},

			// URL : localhost:8080/?Query=mutation+_{addMoviesGenres(movies_id:1, genres_id:1){id, movies_id, genres_id}}
			"addMoviesGenres": &graphql.Field{
				Type:    types.MoviesGenresType,   // Define MoviesGenresType
				Args:    args.AddMoviesGenres,     // Define AddMoviesGenres Arguments (params to add)
				Resolve: resolver.AddMoviesGenres, // Call AddMoviesGenres resolver function
			},

			// URL : localhost:8080/?Query=mutation+_{addReview(review:"Keren Suka saya bos", rate:7, movies_id:7, users_id:2){id, review, rate, users_id, movies_id}}
			"addReview": &graphql.Field{
				Type:    types.ReviewType,   // Define ReviewType
				Args:    args.AddReview,     // Define AddReview Arguments (params to add)
				Resolve: resolver.AddReview, // Call AddReview resolver function
			},
			// // update data by id
			// "updateProduct": &graphql.Field{
			// 	Type:    types.ProductType(),      // Define ProductType
			// 	Args:    args.UpdateProductArgs(), // Define Update Product Arguments (params to update)
			// 	Resolve: resolver.UpdateProduct,   // Call UpdateProduct resolver function
			// },
		},
	},
)
