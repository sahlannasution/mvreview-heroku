package main

import (
	"net/http"

	"github.com/sahlannasution/mvreview-heroku/config"
	"github.com/sahlannasution/mvreview-heroku/middlewares"
	"github.com/sahlannasution/mvreview-heroku/migrator"
	"github.com/sahlannasution/mvreview-heroku/routes"
	"github.com/sahlannasution/mvreview-heroku/schema"
	"github.com/sahlannasution/mvreview-heroku/seeder"

	"github.com/gin-gonic/gin"
)

func main() {

	dbPG := config.Connection() // db Connection
	StrDB := middlewares.StrDB{DB: dbPG}
	migrator.Migrations(dbPG)       // migrate tables
	seeder.SeederUser(dbPG)         // seed Users Data
	seeder.SeederGenres(dbPG)       // seed Genres Data
	seeder.SeederMovies(dbPG)       // seed Movies Data
	seeder.SeederMoviesGenres(dbPG) // seed MoviesGenres Data
	seeder.SeederReview(dbPG)       // seed Reviews Data

	route := gin.Default()
	/* User Signin */
	route.POST("/signin", StrDB.MiddleWare().LoginHandler)

	// Define route
	route.POST("/", StrDB.MiddleWare().MiddlewareFunc(), func(c *gin.Context) {
		// Struvt Query
		type Query struct {
			Query string `json:"query"`
		}

		var query Query

		c.Bind(&query) // Get query params

		result := routes.ExecuteQuery(query.Query, schema.Schema) // Run Query
		c.JSON(http.StatusOK, result)                             // Send Response
	})
	route.Run()
}
