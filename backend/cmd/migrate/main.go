package main

import (
	"context"
	"log"
	"os"

	"github.com/morning-night-dream/article-share/database"
)

func main() {
	client := database.NewClient(
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PASS"),
	)

	ctx := context.Background()

	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("Failed create schema: %v", err)
	}
}
