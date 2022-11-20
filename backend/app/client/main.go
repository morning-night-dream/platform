package main

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/morning-night-dream/platform/app/db/ent/proto/entpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	conn, err := grpc.Dial("localhost:5555", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("failed connecting to server: %s", err)
	}
	defer conn.Close()

	client := entpb.NewArticleServiceClient(conn)

	ctx := context.Background()
	created, err := client.Create(ctx, &entpb.CreateArticleRequest{
		Article: randomArticle(),
	})
	if err != nil {
		se, _ := status.FromError(err)
		log.Fatalf("failed creating article: status=%s message=%s", se.Code(), se.Message())
	}
	log.Printf("article created with id: %s", string(created.Id))

	get, err := client.List(ctx, &entpb.ListArticleRequest{
		PageSize:  10,
		PageToken: "",
		View:      0,
	})
	if err != nil {
		se, _ := status.FromError(err)
		log.Fatalf("failed retrieving user: status=%s message=%s", se.Code(), se.Message())
	}
	for _, art := range get.ArticleList {
		log.Printf("%+v", art)
	}
}

func randomArticle() *entpb.Article {
	return &entpb.Article{
		Id:          []byte{},
		Title:       "test",
		Url:         "https://example.com",
		Description: "test",
		Thumbnail:   "",
		CreatedAt: &timestamppb.Timestamp{
			Seconds: 0,
			Nanos:   0,
		},
		UpdatedAt: &timestamppb.Timestamp{
			Seconds: 0,
			Nanos:   0,
		},
		DeletedAt: &timestamppb.Timestamp{
			Seconds: 0,
			Nanos:   0,
		},
		Tags:         []*entpb.ArticleTag{},
		ReadArticles: []*entpb.ReadArticle{},
	}
}
