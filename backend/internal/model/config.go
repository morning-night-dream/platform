package model

import (
	"os"
)

type config struct {
	Domain              string
	DSN                 string
	RedisURL            string
	NewRelicAppName     string
	NewRelicLicense     string
	FirebaseSecret      string
	FirebaseAPIEndpoint string
	FirebaseAPIKey      string
	AppCoreURL          string
}

var Config config

func init() {
	Config = config{
		Domain:              os.Getenv("DOMAIN"),
		DSN:                 os.Getenv("DATABASE_URL"),
		RedisURL:            os.Getenv("REDIS_URL"),
		NewRelicAppName:     os.Getenv("NEWRELIC_APP_NAME"),
		NewRelicLicense:     os.Getenv("NEWRELIC_LICENSE"),
		FirebaseSecret:      os.Getenv("FIREBASE_SECRET"),
		FirebaseAPIEndpoint: os.Getenv("FIREBASE_API_ENDPOINT"),
		FirebaseAPIKey:      os.Getenv("FIREBASE_API_KEY"),
		AppCoreURL:          os.Getenv("APP_CORE_URL"),
	}
}
