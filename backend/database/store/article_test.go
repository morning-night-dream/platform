package store_test

import (
	"context"
	"testing"

	_ "github.com/mattn/go-sqlite3"

	"github.com/morning-night-dream/article-share/database/ent"
	"github.com/morning-night-dream/article-share/database/ent/enttest"
	"github.com/morning-night-dream/article-share/database/ent/migrate"
	"github.com/morning-night-dream/article-share/database/store"
	"github.com/morning-night-dream/article-share/model"
)

func TestArticleStoreSave(t *testing.T) {
	t.Parallel()

	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
		enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)),
	}

	db := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1", opts...)

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
