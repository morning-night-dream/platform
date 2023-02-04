package main

import (
	"context"
	"log"
	"os"

	"github.com/morning-night-dream/platform/internal/driver/database"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")

	client := database.NewClient(dsn)

	ctx := context.Background()

	if err := client.Debug().Schema.Create(ctx); err != nil {
		log.Fatalf("Failed create schema: %v", err)
	}
}
