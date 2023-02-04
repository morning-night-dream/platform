package main

import (
	"context"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/morning-night-dream/platform/internal/driver/firebase"
)

func main() {
	client := firebase.NewClient(os.Getenv("FIREBASE_SECRET"), "http://localhost:9091/identitytoolkit.googleapis.com", "emulator")

	ctx := context.Background()

	uid := uuid.NewString()

	email := uid + "@example.com"

	pass := uid

	if err := client.CreateUser(ctx, uid, email, pass); err != nil {
		panic(err)
	}

	res, err := client.Login(ctx, email, pass)
	if err != nil {
		panic(err)
	}

	log.Printf("%+v", res)
}
