package store_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/morning-night-dream/article-share/database"
	"github.com/morning-night-dream/article-share/database/store"
	"github.com/morning-night-dream/article-share/model"
)

func setup(ctx context.Context) {
	dsn := fmt.Sprintf("postgres://test:test@%s:54321/test?sslmode=disable", os.Getenv("POSTGRES_HOST"))

	client := database.NewClient(dsn)

	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("Failed create schema: %v", err)
	}
}

func teardown() {
	dsn := fmt.Sprintf("postgres://test:test@%s:54321/test?sslmode=disable", os.Getenv("POSTGRES_HOST"))

	client := database.NewClient(dsn)

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

	dsn := fmt.Sprintf("postgres://test:test@%s:54321/test?sslmode=disable", os.Getenv("POSTGRES_HOST"))

	db := database.NewClient(dsn)

	sa := store.NewArticle(db)

	t.Run("", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()

		if err := sa.Save(ctx, model.Article{
			Title:       "title",
			URL:         "url",
			Description: "description",
			ImageURL:    "image",
		}); err != nil {
			t.Error(err)
		}

		if err := sa.Save(ctx, model.Article{
			Title:       "title",
			URL:         "url",
			Description: "description",
			ImageURL:    "image",
		}); err != nil {
			t.Error(err)
		}
	})
}
