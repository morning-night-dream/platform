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

	"github.com/morning-night-dream/platform/internal/model"
)

const (
	shutdownTime      = 10
	readHeaderTimeout = 30 * time.Second
)

type HTTPServer struct {
	*http.Server
}

func NewHTTPServer(
	handler http.Handler,
) *HTTPServer {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	s := &http.Server{
		Addr:              fmt.Sprintf(":%s", port),
		Handler:           handler,
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
