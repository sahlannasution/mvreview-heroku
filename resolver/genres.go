package resolver

import (
	"mvreview-heroku/config"
	"mvreview-heroku/models"

	"github.com/graphql-go/graphql"
)

// func GetGenresList
func GetGenresList(params graphql.ResolveParams) (interface{}, error) {
	dbPG := config.Connection()
	// genres struct
	var (
		genres []models.Genres
	)
	dbPG.Find(&genres)
	return genres, nil //return genres data response
}

// func AddGenre
func AddGenre(params graphql.ResolveParams) (interface{}, error) {
	name, ok := params.Args["name"].(string) // Get name from params arguments
	dbPG := config.Connection()
	// genres struct
	var (
		genres models.Genres
	)

	if ok {
		genres.Name = name
		dbPG.Create(&genres)
		return genres, nil //return genres data response
	}

	return nil, nil //return genres data response
}
