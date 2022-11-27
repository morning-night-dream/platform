package main

import (
	"os"

	"github.com/morning-night-dream/platform/app/core/database"
	"github.com/morning-night-dream/platform/app/core/database/store"
	"github.com/morning-night-dream/platform/app/core/handler"
	"github.com/morning-night-dream/platform/app/core/server"
)

func main() {
	db := database.NewClient(os.Getenv("DATABASE_URL"))

	sa := store.NewArticle(db)

	ah := handler.NewArticle(*sa)

	hh := handler.NewHealth()

	aua := store.NewAuth(db)

	auh := handler.NewAuth(*aua)

	srv := server.NewHTTPServer(hh, ah, auh)

	srv.Run()
}
