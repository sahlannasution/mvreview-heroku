package resolver

import (
	"mvreview-heroku/config"
	"mvreview-heroku/models"

	"github.com/graphql-go/graphql"
)

// func GetMoviesList
func GetMoviesList(params graphql.ResolveParams) (interface{}, error) {
	dbPG := config.Connection()
	// movies struct
	var (
		movies []models.Movies
	)
	dbPG.Preload("Genres").Find(&movies)
	return movies, nil //return movies data response
}

// func AddMovies
func AddMovies(params graphql.ResolveParams) (interface{}, error) {
	title, checkTitle := params.Args["title"].(string)    // Get title from params arguments
	year, checkYear := params.Args["year"].(int)          // Get year from params arguments
	ratings, checkRatings := params.Args["ratings"].(int) // Get ratings from params arguments
	dbPG := config.Connection()
	// movies struct
	var (
		movies models.Movies
	)

	if checkTitle && checkYear && checkRatings {
		movies.Title = title
		movies.Year = year
		movies.Ratings = ratings
		dbPG.Create(&movies)
		return movies, nil //return movies data response
	}

	return nil, nil //return movies data response
}
