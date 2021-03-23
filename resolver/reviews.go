package resolver

import (
	"github.com/sahlannasution/mvreview-heroku/config"
	"github.com/sahlannasution/mvreview-heroku/models"

	"github.com/graphql-go/graphql"
)

// func GetAllUser
func GetReviewByMovie(params graphql.ResolveParams) (interface{}, error) {
	movies_id, ok := params.Args["movies_id"].(int) // Get movies_id from params arguments
	dbPG := config.Connection()

	var (
		review []models.Reviews
	)
	if ok {
		dbPG.Where("movies_id = ?", movies_id).Find(&review)
		return review, nil //return review data response
	}
	return nil, nil

}

// func GetUser
func GetReviewDetail(params graphql.ResolveParams) (interface{}, error) {
	id, ok := params.Args["id"].(int) // Get id from params arguments
	dbPG := config.Connection()

	var (
		review models.Reviews
	)

	if ok {
		dbPG.Where("id = ?", id).First(&review)
		return review, nil //return users data response
	}
	return nil, nil
}

// func GetUser
func AddReview(params graphql.ResolveParams) (interface{}, error) {
	review, checkReview := params.Args["review"].(string)    // Get review from params arguments
	rate, checkRate := params.Args["rate"].(int)             // Get rate from params arguments
	movies_id, checkMovies := params.Args["movies_id"].(int) // Get movies_id from params arguments
	users_id, checkUsers := params.Args["users_id"].(int)    // Get users_id from params arguments
	dbPG := config.Connection()
	// movies struct
	var (
		reviews models.Reviews
	)

	if checkReview && checkRate && checkMovies && checkUsers {
		reviews.Review = review
		reviews.Rate = rate
		reviews.MoviesID = uint(movies_id)
		reviews.UsersID = uint(users_id)
		dbPG.Create(&reviews)
		return reviews, nil //return reviews data response
	}

	return nil, nil //return reviews data response
}
