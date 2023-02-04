package config

import (
	"os"
)

type core struct {
	Domain              string
	DSN                 string
	RedisURL            string
	NewRelicAppName     string
	NewRelicLicense     string
	FirebaseSecret      string
	FirebaseAPIEndpoint string
	FirebaseAPIKey      string
}

var Core core

func init() {
	Core = core{
		Domain:              os.Getenv("DOMAIN"),
		DSN:                 os.Getenv("DATABASE_URL"),
		RedisURL:            os.Getenv("REDIS_URL"),
		NewRelicAppName:     os.Getenv("NEWRELIC_APP_NAME"),
		NewRelicLicense:     os.Getenv("NEWRELIC_LICENSE"),
		FirebaseSecret:      os.Getenv("FIREBASE_SECRET"),
		FirebaseAPIEndpoint: os.Getenv("FIREBASE_API_ENDPOINT"),
		FirebaseAPIKey:      os.Getenv("FIREBASE_API_KEY"),
	}
}
