package main

import (
	"os"

	"github.com/morning-night-dream/platform/internal/core/database"
	"github.com/morning-night-dream/platform/internal/core/database/store"
	"github.com/morning-night-dream/platform/internal/core/handler"
	"github.com/morning-night-dream/platform/internal/core/server"
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
