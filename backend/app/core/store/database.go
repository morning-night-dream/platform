package store

import (
	// postgres driver.
	_ "github.com/lib/pq"
	"github.com/morning-night-dream/platform/pkg/ent"
)

func NewDatabaseClient(dsn string) *ent.Client {
	client, err := ent.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	return client
}
