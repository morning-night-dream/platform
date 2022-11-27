package database

import (
	"context"
	"log"

	// go-sqlite3.
	_ "github.com/mattn/go-sqlite3"
	"github.com/morning-night-dream/platform/pkg/ent"
)

func NewClient(dsn string) *ent.Client {
	client, err := ent.Open("sqlite3", dsn)

	if err := client.Debug().Schema.Create(context.Background()); err != nil {
		log.Fatalf("Failed create schema: %v", err)
	}

	if err != nil {
		panic(err)
	}

	return client
}
