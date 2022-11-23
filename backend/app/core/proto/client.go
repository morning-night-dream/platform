package proto

import (
	"log"
	"os"

	"github.com/morning-night-dream/platform/app/db/ent/proto/entpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	Article     entpb.ArticleServiceClient
	ArticleTag  entpb.ArticleTagServiceClient
	ReadArticle entpb.ReadArticleServiceClient
	User        entpb.UserServiceClient
	Auth        entpb.AuthServiceClient
}

func NewClient() *Client {
	address := os.Getenv("APP_DB_ADDRESS")

	log.Println(address)

	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Println(err)

		panic(err)
	}

	log.Printf("connection")

	return &Client{
		Article:     entpb.NewArticleServiceClient(conn),
		ArticleTag:  entpb.NewArticleTagServiceClient(conn),
		ReadArticle: entpb.NewReadArticleServiceClient(conn),
		User:        entpb.NewUserServiceClient(conn),
		Auth:        entpb.NewAuthServiceClient(conn),
	}
}
