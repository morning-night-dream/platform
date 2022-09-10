package model_test

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/morning-night-dream/article-share/database"
	"github.com/morning-night-dream/article-share/model"
)

func setup(ctx context.Context) {
	client := database.NewClient(
		os.Getenv("POSTGRES_HOST"),
		"54321",
		"test",
		"test",
		"test",
	)

	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("Failed create schema: %v", err)
	}
}

func teardown() {
	client := database.NewClient(
		os.Getenv("POSTGRES_HOST"),
		"54321",
		"test",
		"test",
		"test",
	)

	ctx := context.Background()

	if _, err := client.Article.Delete().Exec(ctx); err != nil {
		log.Fatal(err)
	}
}

func TestMain(m *testing.M) {
	ctx := context.Background()

	setup(ctx)

	status := m.Run()

	teardown()

	os.Exit(status)
}

func TestArticleStoreSave(t *testing.T) {
	t.Parallel()

	db := database.NewClient(
		os.Getenv("POSTGRES_HOST"),
		"54321",
		"test",
		"test",
		"test",
	)

	store := model.NewArticleStore(db)

	t.Run("", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()

		if err := store.Save(ctx, model.Article{
			Title:       "title",
			URL:         "url",
			Description: "description",
			ImageURL:    "image",
		}); err != nil {
			t.Error(err)
		}

		if err := store.Save(ctx, model.Article{
			Title:       "title",
			URL:         "url",
			Description: "description",
			ImageURL:    "image",
		}); err != nil {
			t.Error(err)
		}
	})
}
