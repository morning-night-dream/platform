package main

import (
	"context"
	"log"

	"github.com/morning-night-dream/article-share/database"
)

func main() {
	client := database.NewClient()

	ctx := context.Background()

	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("Failed create schema: %v", err)
	}
}
