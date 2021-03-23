package seeder

import (
	"fmt"
	"log"
	"mvreview-heroku/models"
	"strconv"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeederReview(db *gorm.DB) {
	var reviewArray = [...][4]string{
		{"1", "4", "Keren Banget Filmnya", "9"},
		{"1", "5", "Keren Banget Filmnya", "8"},
		{"1", "6", "Bagus jalan cerita nya", "7"},
		{"2", "1", "sad ending", "3"},
		{"2", "2", "bagus tapi kurang nendang", "5"},
		{"2", "5", "siapa nih yang bikin karakter nya bagus banget", "8"},
		{"3", "1", "karakter utama nya gabagsu", "5"},
		{"3", "4", "ok", "6"},
		{"3", "6", "gooood", "7"},
	}

	var review models.Reviews

	for _, data := range reviewArray {
		// Get Data from Array
		user_id, _ := strconv.ParseInt(data[0], 10, 64)
		movie_id, _ := strconv.ParseInt(data[1], 10, 64)
		rate, _ := strconv.ParseInt(data[3], 10, 64)
		review.UsersID = uint(user_id)
		review.MoviesID = uint(movie_id)
		review.Review = data[2]
		review.Rate = int(rate)
		review.ID = 0 // declare id dimulai dari 0, karena auto increment
		db.Create(&review)
	}
	fmt.Println("Review Data has been Seed!")
}

func SeederMoviesGenres(db *gorm.DB) {
	var moviesGenresArray = [...][2]string{
		{"1", "4"},
		{"1", "5"},
		{"1", "6"},
		{"2", "1"},
		{"2", "2"},
		{"2", "7"},
		{"3", "1"},
		{"3", "2"},
		{"3", "7"},
	}

	var moviesGenres models.MoviesGenres

	for _, data := range moviesGenresArray {
		// Get Data from Array
		movie_id, _ := strconv.ParseInt(data[0], 10, 64)
		genre_id, _ := strconv.ParseInt(data[1], 10, 64)
		moviesGenres.MoviesID = uint(movie_id)
		moviesGenres.GenresID = uint(genre_id)
		moviesGenres.ID = 0
		db.Create(&moviesGenres)
	}
	fmt.Println("MoviesGenres Data has been Seed!")
}

func SeederMovies(db *gorm.DB) {
	var moviesArray = [...][3]string{
		{"Parasite", "2019", "0"},
		{"The Avengers", "2017", "0"},
		{"Spiderman", "2010", "0"},
		{"Doraemon Stand By Me 2", "2020", "0"},
		{"Parasyte", "2016", "0"},
		{"Harry Potter", "2014", "0"},
		{"Ayat-Ayat Cinta", "2006", "0"},
	}

	var movies models.Movies

	for _, data := range moviesArray {
		// Get Data from Array
		year, _ := strconv.ParseInt(data[1], 10, 64)
		rating, _ := strconv.ParseInt(data[2], 10, 64)
		movies.Title = data[0]
		movies.Year = int(year)
		movies.Ratings = int(rating)
		movies.ID = 0 // declare id dimulai dari 0, karena auto increment
		db.Create(&movies)
	}
	fmt.Println("Movies Data has been Seed!")
}

func SeederGenres(db *gorm.DB) {
	var genresArray = [...]string{
		"Action",
		"Sci-Fi",
		"Mystery",
		"Comedy",
		"Drama",
		"Thriller",
		"Adventure",
		"Animation",
	}

	var genres models.Genres

	for _, data := range genresArray {
		// Get Data from Array
		genres.Name = data
		genres.ID = 0
		db.Create(&genres)
	}
	fmt.Println("Genre Data has been Seed!")
}

func SeederUser(db *gorm.DB) {
	var userArray = [...][4]string{
		{"Sahlan.Nasution@xapiens.id", "admin@1234", "Sahlan Admin Xapiens", "admin"},
		// {"sahlan.nasution07@gmail.com", "guest@1234", "Sahlan Nasution", "guest"},
		// {"zahland.nasution@gmail.com", "guest@1234", "Sahlan NST", "guest"},
		{"armanzulfikri@mail.com", "guest@1234", "Arman Zulfikri", "guest"},
		{"wawan_setiawan@mail.com", "guest@1234", "Wawan Setiawan", "guest"},
		{"mas.dimas@mail.com", "guest@1234", "Mas Dimas", "guest"},
		{"dimas.maskun@mail.com", "guest@1234", "Dimas Chan wkwk", "guest"},
	}

	var user models.Users

	for _, data := range userArray {
		// Get Data from Array
		user.Email = data[0]
		user.Password = data[1]
		user.FullName = data[2]
		user.Role = data[3]

		// Encrypt Password using Bcrypt
		encrypt, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

		// Check if error while encrypting
		if err != nil {
			log.Println(err)
		}

		user.Password = string(encrypt)
		user.ID = 0 // declare id dimulai dari 0, karena auto increment
		db.Create(&user)
	}
	fmt.Println("Users Data has been Seed!")
}
