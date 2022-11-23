package main

import (
	"context"
	"log"
	"net"
	"os"

	_ "github.com/lib/pq"
	"github.com/morning-night-dream/platform/app/db/ent"
	"github.com/morning-night-dream/platform/app/db/ent/proto/entpb"
	"google.golang.org/grpc"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")

	client, err := ent.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	client = client.Debug()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	server := grpc.NewServer()

	entpb.RegisterArticleServiceServer(server, entpb.NewArticleService(client))
	entpb.RegisterArticleTagServiceServer(server, entpb.NewArticleTagService(client))
	entpb.RegisterReadArticleServiceServer(server, entpb.NewReadArticleService(client))
	entpb.RegisterUserServiceServer(server, entpb.NewUserService(client))
	entpb.RegisterAuthServiceServer(server, entpb.NewAuthService(client))

	log.Println("db app run ...")

	lis, err := net.Listen("tcp", ":5555")
	if err != nil {
		log.Fatalf("failed listening: %s", err)
	}

	if err := server.Serve(lis); err != nil {
		log.Fatalf("server ended: %s", err)
	}
}
