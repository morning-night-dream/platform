package database

import (
	// postgres driver.
	_ "github.com/lib/pq"
	"github.com/morning-night-dream/platform/pkg/ent"
)

func NewClient(dsn string) *ent.Client {
	client, err := ent.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	return client
}
