package main

import (
	"os"

	"github.com/morning-night-dream/platform/internal/cache"
	"github.com/morning-night-dream/platform/internal/database"
	"github.com/morning-night-dream/platform/internal/database/store"
	"github.com/morning-night-dream/platform/internal/firebase"
	"github.com/morning-night-dream/platform/internal/handler"
	"github.com/morning-night-dream/platform/internal/model"
	"github.com/morning-night-dream/platform/internal/server"
)

func main() {
	db := database.NewClient(os.Getenv("DATABASE_URL"))

	cache := cache.NewClient()

	sa := store.NewArticle(db)

	fb := firebase.NewClient(model.Config.FirebaseSecret, model.Config.FirebaseAPIEndpoint, model.Config.FirebaseAPIKey)

	handle := handler.NewHandle(fb, cache)

	ah := handler.NewArticle(sa, handle)

	hh := handler.NewHealth()

	auh := handler.NewAuth(handle)

	srv := server.NewHTTPServer(hh, ah, auh)

	srv.Run()
}
