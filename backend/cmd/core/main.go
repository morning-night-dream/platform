package main

import (
	"github.com/morning-night-dream/platform/internal/cache"
	"github.com/morning-night-dream/platform/internal/database"
	"github.com/morning-night-dream/platform/internal/firebase"
	"github.com/morning-night-dream/platform/internal/handler"
	"github.com/morning-night-dream/platform/internal/model"
	"github.com/morning-night-dream/platform/internal/server"
)

func main() {
	db := database.NewClient(model.Config.DSN)

	cache := cache.NewClient(model.Config.CacheURL)

	da := database.NewArticle(db)

	fb := firebase.NewClient(model.Config.FirebaseSecret, model.Config.FirebaseAPIEndpoint, model.Config.FirebaseAPIKey)

	handle := handler.NewHandle(fb, cache)

	ah := handler.NewArticle(da, handle)

	hh := handler.NewHealth()

	auh := handler.NewAuth(handle)

	srv := server.NewHTTPServer(hh, ah, auh)

	srv.Run()
}
