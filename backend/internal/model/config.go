package model

import (
	"os"
)

type config struct {
	Domain              string
	NewRelicAppName     string
	NewRelicLicense     string
	FirebaseSecret      string
	FirebaseAPIEndpoint string
	FirebaseAPIKey      string
}

var Config config

func init() {
	Config = config{
		Domain:              os.Getenv("DOMAIN"),
		NewRelicAppName:     os.Getenv("NEWRELIC_APP_NAME"),
		NewRelicLicense:     os.Getenv("NEWRELIC_LICENSE"),
		FirebaseSecret:      os.Getenv("FIREBASE_SECRET"),
		FirebaseAPIEndpoint: os.Getenv("FIREBASE_API_ENDPOINT"),
		FirebaseAPIKey:      os.Getenv("FIREBASE_API_KEY"),
	}
}