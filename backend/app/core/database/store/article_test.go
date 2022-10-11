package store_test

import (
	"context"
	"reflect"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/morning-night-dream/article-share/app/core/database/store"
	"github.com/morning-night-dream/article-share/app/core/model"
	"github.com/morning-night-dream/article-share/pkg/ent"
	"github.com/morning-night-dream/article-share/pkg/ent/enttest"
	"github.com/morning-night-dream/article-share/pkg/ent/migrate"
)

func TestArticleStoreSave(t *testing.T) {
	t.Parallel()

	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
		enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)),
	}

	db := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1", opts...)

	sa := store.NewArticle(db)

	t.Run("記事を保存できる", func(t *testing.T) {
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

	t.Run("記事を取得できる", func(t *testing.T) {
		t.Parallel()

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
}
