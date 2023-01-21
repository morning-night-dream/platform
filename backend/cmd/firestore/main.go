package main

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"github.com/google/uuid"
)

func main() {
	os.Setenv("FIRESTORE_EMULATOR_HOST", "localhost:9100")

	ctx := context.Background()

	conf := &firebase.Config{
		ProjectID: "emulator",
	}

	app, err := firebase.NewApp(ctx, conf)

	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	defer client.Close()

	path := "session"

	id := uuid.NewString()

	// データ保存
	if _, err := client.Collection(path).Doc(id).Create(ctx, map[string]interface{}{
		"key": "value",
	}); err != nil {
		log.Fatalln(err)
	}

	// データ取得
	res, err := client.Collection(path).Doc(id).Get(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%+v", res.Data())

	// データ削除
	if _, err := client.Collection(path).Doc(id).Delete(ctx); err != nil {
		log.Fatalln(err)
	}

	// データ取得
	res, err = client.Collection(path).Doc(id).Get(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%+v", res.Data())

}
