package main

import (
	"github.com/morning-night-dream/platform/internal/adapter/gateway"
	"github.com/morning-night-dream/platform/internal/adapter/handler"
	"github.com/morning-night-dream/platform/internal/driver/config"
	"github.com/morning-night-dream/platform/internal/driver/database"
	"github.com/morning-night-dream/platform/internal/driver/firebase"
	"github.com/morning-night-dream/platform/internal/driver/redis"
	"github.com/morning-night-dream/platform/internal/driver/server"
)

func main() {
	db := database.NewClient(config.Core.DSN)

	cache := redis.NewClient(config.Core.RedisURL)

	da := gateway.NewArticle(db)

	fb := firebase.NewClient(config.Core.FirebaseSecret, config.Core.FirebaseAPIEndpoint, config.Core.FirebaseAPIKey)

	handle := handler.NewHandle(fb, cache)

	ah := handler.NewArticle(da, handle)

	hh := handler.NewHealth()

	auh := handler.NewAuth(handle)

	ch := server.NewConnectHandler(hh, ah, auh)

	srv := server.NewHTTPServer(ch)

	srv.Run()
}
