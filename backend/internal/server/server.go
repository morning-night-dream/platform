package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bufbuild/connect-go"
	"github.com/morning-night-dream/platform/internal/handler"
	"github.com/morning-night-dream/platform/internal/model"
	"github.com/morning-night-dream/platform/pkg/proto/article/v1/articlev1connect"
	"github.com/morning-night-dream/platform/pkg/proto/auth/v1/authv1connect"
	"github.com/morning-night-dream/platform/pkg/proto/health/v1/healthv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

const (
	shutdownTime      = 10
	readHeaderTimeout = 30 * time.Second
)

type HTTPServer struct {
	*http.Server
}

func NewHTTPServer(
	health *handler.Health,
	article *handler.Article,
	auth *handler.Auth,
) *HTTPServer {
	interceptor := connect.WithInterceptors(NewInterceptor())

	mux := NewRouter(
		NewRoute(healthv1connect.NewHealthServiceHandler(health, interceptor)),
		NewRoute(articlev1connect.NewArticleServiceHandler(article, interceptor)),
		NewRoute(authv1connect.NewAuthServiceHandler(auth, interceptor)),
	).Mux()

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	s := &http.Server{
		Addr:              fmt.Sprintf(":%s", port),
		Handler:           h2c.NewHandler(mux, &http2.Server{}),
		ReadHeaderTimeout: readHeaderTimeout,
	}

	return &HTTPServer{
		Server: s,
	}
}

func (s *HTTPServer) Run() {
	log.Printf("env is %s\n", model.Env.String())
	log.Printf("Server running on %s", s.Addr)

	go func() {
		if err := s.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Printf("Server closed with error: %s", err.Error())

			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGTERM, os.Interrupt)

	log.Printf("SIGNAL %d received, then shutting down...\n", <-quit)

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTime*time.Second)

	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Printf("Failed to gracefully shutdown: %d", err)
	}

	log.Printf("HTTPServer shutdown")
}
