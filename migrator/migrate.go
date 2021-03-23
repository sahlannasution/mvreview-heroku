package migrator

import (
	"fmt"

	"github.com/sahlannasution/mvreview-heroku/models"

	"gorm.io/gorm"
)

func Migrations(db *gorm.DB) {

	// Drop Table Users
	db.Migrator().DropTable(&models.Users{})
	if check := db.Migrator().HasTable(&models.Users{}); !check {
		db.Migrator().CreateTable(&models.Users{})
		fmt.Println("Table Users has been created!")
	}

	// Drop Table Genres
	db.Migrator().DropTable(&models.Genres{})
	if check := db.Migrator().HasTable(&models.Genres{}); !check {
		db.Migrator().CreateTable(&models.Genres{})
		fmt.Println("Table Genres has been created!")
	}

	// Drop Table Movies
	db.Migrator().DropTable(&models.Movies{})
	if check := db.Migrator().HasTable(&models.Movies{}); !check {
		db.Migrator().CreateTable(&models.Movies{})
		fmt.Println("Table Movies has been created!")
	}

	// Drop Table MoviesGenres
	db.Migrator().DropTable(&models.MoviesGenres{})
	if check := db.Migrator().HasTable(&models.MoviesGenres{}); !check {
		db.Migrator().CreateTable(&models.MoviesGenres{})
		fmt.Println("Table MoviesGenres has been created!")
	}

	// Drop Table MoviesGenres
	db.Migrator().DropTable(&models.Reviews{})
	if check := db.Migrator().HasTable(&models.Reviews{}); !check {
		db.Migrator().CreateTable(&models.Reviews{})
		fmt.Println("Table Reviews has been created!")
	}
}
