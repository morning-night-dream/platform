package helper

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/morning-night-dream/platform/internal/driver/database"
	"github.com/morning-night-dream/platform/pkg/ent"
)

func BulkInsert(t *testing.T, count int) {
	t.Helper()

	dsn := os.Getenv("DATABASE_URL")

	client := database.NewClient(dsn)

	defer client.Close()

	ids := make([]string, count)

	for i := 0; i < count; i++ {
		ids[i] = fmt.Sprintf("00000000-0000-0000-0000-000000000%03d", i)
	}

	bulk := make([]*ent.ArticleCreate, count)

	for i, id := range ids {
		bulk[i] = client.Article.Create().
			SetID(uuid.MustParse(id)).
			SetTitle("title-" + id).
			SetURL("https://example.com/" + id).
			SetDescription("description").
			SetThumbnail("https://example.com/" + id).
			SetCreatedAt(time.Now()).
			SetUpdatedAt(time.Now())
	}

	if err := client.Article.CreateBulk(bulk...).OnConflict().UpdateNewValues().DoNothing().Exec(context.Background()); err != nil {
		t.Fatal(err)
	}
}

func BulkDelete(t *testing.T, count int) {
	t.Helper()

	dsn := os.Getenv("DATABASE_URL")

	client := database.NewClient(dsn)

	defer client.Close()

	ids := make([]string, count)

	for i := 0; i < count; i++ {
		ids[i] = fmt.Sprintf("00000000-0000-0000-0000-000000000%03d", i)
	}

	tx, err := client.Tx(context.Background())
	if err != nil {
		t.Error(err)

		return
	}

	for _, id := range ids {
		if err := tx.Article.DeleteOneID(uuid.MustParse(id)).Exec(context.Background()); err != nil {
			t.Error(err)

			_ = tx.Rollback()
		}
	}

	if err := tx.Commit(); err != nil {
		t.Error(err)
	}
}

func DeleteOne(t *testing.T, id string) {
	t.Helper()

	dsn := os.Getenv("DATABASE_URL")

	client := database.NewClient(dsn)

	defer client.Close()

	tx, err := client.Tx(context.Background())
	if err != nil {
		t.Error(err)

		return
	}

	if err := tx.Article.DeleteOneID(uuid.MustParse(id)).Exec(context.Background()); err != nil {
		t.Error(err)

		_ = tx.Rollback()
	}

	if err := tx.Commit(); err != nil {
		t.Error(err)
	}
}
