package main

import (
	"os"

	"github.com/morning-night-dream/platform/app/core/database"
	"github.com/morning-night-dream/platform/app/core/database/store"
	"github.com/morning-night-dream/platform/app/core/handler"
	"github.com/morning-night-dream/platform/app/core/proto"
	"github.com/morning-night-dream/platform/app/core/server"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")

	secret := os.Getenv("SLACK_SIGNING_SECRET")

	db := database.NewClient(dsn)

	sa := store.NewArticle(db)

	pb := proto.NewClient()

	ah := handler.NewArticle(*sa, pb)

	sh := handler.NewSlack(secret, sa)

	aua := store.NewAuth(db)

	auh := handler.NewAuth(*aua)

	srv := server.NewHTTPServer(ah, auh, sh)

	srv.Run()
}
