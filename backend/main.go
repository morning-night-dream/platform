package main

import (
	"log"
	"net/http"
	"os"
	"time"

	articlev1connect "github.com/morning-night-dream/article-share/api/article/v1/articlev1connect"
	"github.com/morning-night-dream/article-share/database"
	"github.com/morning-night-dream/article-share/database/store"
	"github.com/morning-night-dream/article-share/handler"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

const timeout = 10

func main() {
	dsn := os.Getenv("DATABASE_URL")

	secret := os.Getenv("SLACK_SIGNING_SECRET")

	db := database.NewClient(dsn)

	sa := store.NewArticle(db)

	ah := handler.NewArticleHandler(*sa)

	sh := handler.NewSlackHandler(secret, sa)

	mux := http.NewServeMux()

	path, handler := articlev1connect.NewArticleServiceHandler(ah)

	mux.Handle(path, handler)

	mux.HandleFunc("/api/slack/events", sh.Events)

	port := ":8080"

	server := &http.Server{
		Addr:              port,
		Handler:           h2c.NewHandler(mux, &http2.Server{}),
		ReadHeaderTimeout: timeout * time.Second,
	}

	log.Printf("start receiving at %s\n", port)
	log.Fatal(server.ListenAndServe())
}
