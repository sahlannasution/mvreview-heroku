package logger

import (
	"log"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/joho/godotenv"
)

func Sentry(paramError error) {
	var dsnSentry string

	if err := godotenv.Load(".env"); err != nil {
		// fmt.Println("This is the Error : ", err)
		sentry.CaptureException(err)
		log.Fatalf(err.Error())
	} else {
		dsnSentry = os.Getenv("DSN_SENTRY")
	}
	err := sentry.Init(sentry.ClientOptions{
		Dsn: dsnSentry,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	defer sentry.Flush(2 * time.Second)

	sentry.CaptureException(paramError)
}

func SentryStr(message string) {
	var dsnSentry string

	if err := godotenv.Load(".env"); err != nil {
		// fmt.Println("This is the Error : ", err)
		sentry.CaptureException(err)
		log.Fatalf(err.Error())
	} else {
		dsnSentry = os.Getenv("DSN_SENTRY")
	}
	err := sentry.Init(sentry.ClientOptions{
		Dsn: dsnSentry,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	defer sentry.Flush(2 * time.Second)

	sentry.CaptureMessage(message)
}
