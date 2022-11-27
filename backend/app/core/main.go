package main

import (
	"github.com/morning-night-dream/platform/app/core/database"
	"github.com/morning-night-dream/platform/app/core/database/store"
	"github.com/morning-night-dream/platform/app/core/handler"
	"github.com/morning-night-dream/platform/app/core/server"
)

func main() {
	db := database.NewClient("file:database?mode=memory&cache=shared&_fk=1")

	sa := store.NewArticle(db)

	ah := handler.NewArticle(*sa)

	hh := handler.NewHealth()

	aua := store.NewAuth(db)

	auh := handler.NewAuth(*aua)

	srv := server.NewHTTPServer(hh, ah, auh)

	srv.Run()
}
