package store

import (
	"context"
	"database/sql"
	"log"

	"github.com/morning-night-dream/article-share/app/core/model"
	"github.com/morning-night-dream/article-share/pkg/ent"
	"github.com/pkg/errors"
)

type Article struct {
	db *ent.Client
}

func NewArticle(db *ent.Client) *Article {
	return &Article{
		db: db,
	}
}

func (a Article) Save(ctx context.Context, article model.Article) error {
	err := a.db.Article.Create().
		SetTitle(article.Title).
		SetDescription(article.Description).
		SetURL(article.URL).
		SetImageURL(article.ImageURL).
		OnConflict().
		DoNothing().
		Exec(ctx)
	if err != nil {
		// https://github.com/ent/ent/issues/2176 により、
		// on conflict do nothingとしてもerror no rowsが返るため、個別にハンドリングする
		if errors.Is(err, sql.ErrNoRows) {
			log.Print(err)

			return nil
		}

		return errors.Wrap(err, "failed to save")
	}

	return nil
}
