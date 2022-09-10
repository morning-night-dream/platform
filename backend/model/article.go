package model

import (
	"context"
	"database/sql"
	"log"

	"github.com/morning-night-dream/article-share/ent"
	"github.com/pkg/errors"
)

type Article struct {
	ID          string `json:"id"`
	URL         string `json:"url"`
	Title       string `json:"title"`
	ImageURL    string `json:"imageUrl"`
	Description string `json:"description"`
}

type ArticleStore struct {
	db *ent.Client
}

func NewArticleStore(db *ent.Client) *ArticleStore {
	return &ArticleStore{
		db: db,
	}
}

func (a ArticleStore) Save(ctx context.Context, article Article) error {
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
