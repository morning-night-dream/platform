package helper

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/morning-night-dream/platform/internal/driver/database"
	"github.com/morning-night-dream/platform/pkg/ent"
)

type ArticleDB struct {
	T      *testing.T
	client *ent.Client
}

func NewArticleDB(
	t *testing.T,
	dsn string,
) *ArticleDB {
	t.Helper()

	return &ArticleDB{
		T:      t,
		client: database.NewClient(dsn),
	}
}

func (adb *ArticleDB) Close() error {
	return adb.client.Close()
}

func (adb *ArticleDB) BulkInsert(ids []string) {
	adb.T.Helper()

	bulk := make([]*ent.ArticleCreate, len(ids))

	for i, id := range ids {
		bulk[i] = adb.client.Article.Create().
			SetID(uuid.MustParse(id)).
			SetTitle("title-" + id).
			SetURL("https://example.com/" + id).
			SetDescription("description").
			SetThumbnail("https://example.com/" + id).
			SetCreatedAt(time.Now()).
			SetUpdatedAt(time.Now())
	}

	if err := adb.client.Article.CreateBulk(bulk...).OnConflict().UpdateNewValues().DoNothing().Exec(context.Background()); err != nil {
		adb.T.Fatal(err)
	}
}

func (adb ArticleDB) BulkDelete(ids []string) {
	adb.T.Helper()

	tx, err := adb.client.Tx(context.Background())
	if err != nil {
		adb.T.Error(err)

		return
	}

	for _, id := range ids {
		if err := tx.Article.DeleteOneID(uuid.MustParse(id)).Exec(context.Background()); err != nil {
			adb.T.Error(err)

			_ = tx.Rollback()
		}
	}

	if err := tx.Commit(); err != nil {
		adb.T.Error(err)
	}
}
