package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/morning-night-dream/platform/internal/client"
	"github.com/morning-night-dream/platform/internal/controller"
	"github.com/morning-night-dream/platform/internal/model"
	"github.com/morning-night-dream/platform/internal/server"
	"github.com/morning-night-dream/platform/pkg/openapi"
)

func main() {
	c, err := client.New().Of(model.Config.AppCoreURL)
	if err != nil {
		panic(err)
	}

	ctr := controller.New(c)

	handler := openapi.HandlerFromMuxWithBaseURL(ctr, chi.NewRouter(), "/api")

	srv := server.NewHTTPServer(handler)

	srv.Run()
}
