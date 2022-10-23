package main

import (
	"os"

	"github.com/morning-night-dream/platform/app/core/handler"
	"github.com/morning-night-dream/platform/app/core/server"
	"github.com/morning-night-dream/platform/app/core/store"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")

	secret := os.Getenv("SLACK_SIGNING_SECRET")

	db := store.NewDatabaseClient(dsn)

	firebase := store.NewFirebaseClient()

	sa := store.NewArticle(db)

	ah := handler.NewArticle(*sa)

	sh := handler.NewSlack(secret, sa)

	aua := store.NewAuth(db, *firebase)

	auh := handler.NewAuth(*aua)

	srv := server.NewHTTPServer(ah, auh, sh)

	srv.Run()
}
