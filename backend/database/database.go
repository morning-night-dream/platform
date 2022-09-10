package database

import (
	"fmt"
	"os"

	// postgres driver.
	_ "github.com/lib/pq"
	"github.com/morning-night-dream/article-share/ent"
)

func NewClient() *ent.Client {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	name := os.Getenv("POSTGRES_DB")
	pass := os.Getenv("POSTGRES_PASS")

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
