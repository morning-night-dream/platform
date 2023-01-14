package helper

import (
	"os"
	"testing"
)

func GetEMail(t *testing.T) string {
	t.Helper()

	return os.Getenv("TEST_USER_EMAIL")
}

func GetPassword(t *testing.T) string {
	t.Helper()

	return os.Getenv("TEST_USER_PASSWORD")
}
