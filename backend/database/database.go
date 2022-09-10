package database

import (
	"fmt"

	// postgres driver.
	_ "github.com/lib/pq"
	"github.com/morning-night-dream/article-share/ent"
)

func NewClient(
	host string,
	port string,
	user string,
	name string,
	pass string,
) *ent.Client {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		user,
		name,
		pass,
	)

	client, err := ent.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	return client
}
