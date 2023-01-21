package helper

import (
	"os"
	"testing"
)

var apiKey = ""

func GetAPIKey(t *testing.T) string {
	t.Helper()

	if apiKey != "" {
		return apiKey
	}

	k := os.Getenv("API_KEY")

	if k == "" {
		k = "e2e"
	}

	apiKey = k

	return apiKey
}

var endpoint = ""

func GetEndpoint(t *testing.T) string {
	t.Helper()

	if endpoint != "" {
		return endpoint
	}

	ep := os.Getenv("ENDPOINT")

	if ep == "" {
		ep = "http://localhost:8081"
	}

	endpoint = ep

	return endpoint
}

var email = ""

func GetEMail(t *testing.T) string {
	t.Helper()

	if email != "" {
		return email
	}

	email = os.Getenv("TEST_USER_EMAIL")

	return email
}

var password = ""

func GetPassword(t *testing.T) string {
	t.Helper()

	if password != "" {
		return password
	}

	password = os.Getenv("TEST_USER_PASSWORD")

	return password
}
