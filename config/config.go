package config // import package

import (
	"fmt"
	"log"
	"os"

	logger "github.com/sahlannasution/mvreview-heroku/log"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connection func
func Connection() *gorm.DB {
	var userDB, passDB, hostDB, portDB, namaDB, ssl, timeZone string
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf(err.Error())
	} else {
		userDB = os.Getenv("DB_USER")
		passDB = os.Getenv("DB_PASS")
		hostDB = os.Getenv("DB_HOST")
		portDB = os.Getenv("DB_PORT")
		namaDB = os.Getenv("DB_NAME")
		ssl = os.Getenv("DB_SSLMODE")
		timeZone = os.Getenv("DB_TIMEZONE")
	}

	conn :=
		" host=" + hostDB +
			" user=" + userDB +
			" password=" + passDB +
			" dbname=" + namaDB +
			" port=" + portDB +
			" sslmode=" + ssl +
			" TimeZone=" + timeZone

	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})

	if err != nil {
		logger.Sentry(err)
		panic("Error connect to database!")

	} else {
		fmt.Println("Sucessfully connected to database!")
	}

	return db
}
