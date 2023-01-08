package model

import "os"

type config struct {
	NewRelicAppName string
	NewRelicLicense string
}

var Config config

func init() {
	Config = config{
		NewRelicAppName: os.Getenv("NEWRELIC_APP_NAME"),
		NewRelicLicense: os.Getenv("NEWRELIC_LICENSE"),
	}
}
