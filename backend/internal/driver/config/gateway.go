package config

import (
	"os"
)

type gateway struct {
	AppCoreURL string
}

var Gateway gateway

func init() {
	Gateway = gateway{
		AppCoreURL: os.Getenv("APP_CORE_URL"),
	}
}
