package main

import (
	"log"
	"net/http"
	"time"

	articlev1connect "github.com/morning-night-dream/article-share/api/article/v1/articlev1connect"
	"github.com/morning-night-dream/article-share/database"
	"github.com/morning-night-dream/article-share/handler"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

const timeout = 10

func main() {
	db := database.NewClient()

	ah := handler.NewArticleHandler(db)

	mux := http.NewServeMux()

	path, handler := articlev1connect.NewArticleServiceHandler(ah)

	mux.Handle(path, handler)

	port := ":8080"

	server := &http.Server{
		Addr:              port,
		Handler:           h2c.NewHandler(mux, &http2.Server{}),
		ReadHeaderTimeout: timeout * time.Second,
	}

	log.Printf("start receiving at %s\n", port)
	log.Fatal(server.ListenAndServe())
}
