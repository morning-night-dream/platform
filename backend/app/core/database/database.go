package database

import (
	"context"
	"log"

	// postgres driver.
	_ "github.com/lib/pq"
	"github.com/morning-night-dream/platform/pkg/ent"
)

func NewClient(dsn string) *ent.Client {
	client, err := ent.Open("postgres", dsn)

	if err := client.Debug().Schema.Create(context.Background()); err != nil {
		log.Fatalf("Failed create schema: %v", err)
	}

	if err != nil {
		panic(err)
	}

	return client
}
