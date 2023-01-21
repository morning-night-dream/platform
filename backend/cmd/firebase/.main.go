package main

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/google/uuid"
	"google.golang.org/api/option"
)

func main() {
	opt := option.WithCredentialsJSON([]byte(os.Getenv("FIREBASE_SECRET")))

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v", err)
	}

	client, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error create auth client: %v", err)
	}

	ctx := context.Background()

	uid := uuid.NewString()

	params := (&auth.UserToCreate{}).
		UID(uid).
		Email(uid + "@example.com").
		EmailVerified(false).
		Password(uid).
		Disabled(false)

	res, err := client.CreateUser(ctx, params)
	if err != nil {
		panic(err)
	}

	log.Printf("created user: %+v", res)

	log.Printf("%+v", res.UID)
}
