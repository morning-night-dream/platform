package store_test

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"github.com/morning-night-dream/platform/app/core/database/store"
	"github.com/morning-night-dream/platform/app/core/model"
	"github.com/morning-night-dream/platform/pkg/ent"
	"github.com/morning-night-dream/platform/pkg/ent/enttest"
	"github.com/morning-night-dream/platform/pkg/ent/migrate"
)

func TestArticleStoreSave(t *testing.T) {
	t.Parallel()

	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
		enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)),
	}

	t.Run("記事を保存できる", func(t *testing.T) {
		t.Parallel()

		dsn := fmt.Sprintf("file:%s?mode=memory&cache=shared&_fk=1", uuid.NewString())

		db := enttest.Open(t, "sqlite3", dsn, opts...)

		sa := store.NewArticle(db)

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

	t.Run("記事を取得できる", func(t *testing.T) {
		t.Parallel()

		dsn := fmt.Sprintf("file:%s?mode=memory&cache=shared&_fk=1", uuid.NewString())

		db := enttest.Open(t, "sqlite3", dsn, opts...)

		sa := store.NewArticle(db)

		ctx := context.Background()

		if err := sa.Save(ctx, model.Article{
			Title:       "title1",
			URL:         "url1",
			Description: "description1",
			ImageURL:    "image1",
		}); err != nil {
			t.Error(err)
		}

		if err := sa.Save(ctx, model.Article{
			Title:       "title1",
			URL:         "url1",
			Description: "description1",
			ImageURL:    "image1",
		}); err != nil {
			t.Error(err)
		}

		got, err := sa.FindAll(ctx, 1, 0)
		if err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(1, len(got)) {
			t.Errorf("NewArticle() = %v, want %v", len(got), 1)
		}
	})

	t.Run("記事を論理削除できる", func(t *testing.T) {
		t.Parallel()

		dsn := fmt.Sprintf("file:%s?mode=memory&cache=shared&_fk=1", uuid.NewString())

		db := enttest.Open(t, "sqlite3", dsn, opts...)

		sa := store.NewArticle(db)

		ctx := context.Background()

		if err := sa.Save(ctx, model.Article{
			Title:       "title1",
			URL:         "url1",
			Description: "description1",
			ImageURL:    "image1",
		}); err != nil {
			t.Error(err)
		}

		if err := sa.Save(ctx, model.Article{
			Title:       "title2",
			URL:         "url2",
			Description: "description2",
			ImageURL:    "image2",
		}); err != nil {
			t.Error(err)
		}

		articles, err := sa.FindAll(ctx, 10, 0)
		if err != nil {
			t.Error(err)
		}

		id := articles[0].ID

		if err := sa.LogicalDelete(ctx, id); err != nil {
			t.Error(err)
		}

		got, err := sa.FindAll(ctx, 10, 0)
		if err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(1, len(got)) {
			t.Errorf("NewArticle() = %v, want %v", len(got), 1)
		}
	})
}
