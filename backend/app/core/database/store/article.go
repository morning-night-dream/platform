package store

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/morning-night-dream/platform/app/core/model"
	"github.com/morning-night-dream/platform/pkg/ent"
	"github.com/morning-night-dream/platform/pkg/ent/article"
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
	id, err := uuid.Parse(article.ID)
	if err != nil {
		id = uuid.New()
	}

	err = a.db.Article.Create().
		SetID(id).
		SetTitle(article.Title).
		SetDescription(article.Description).
		SetURL(article.URL).
		SetThumbnail(article.Thumbnail).
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

	if len(article.Tags) == 0 {
		return nil
	}

	bulk := make([]*ent.ArticleTagCreate, len(article.Tags))
	for i, tag := range article.Tags {
		bulk[i] = a.db.ArticleTag.Create().
			SetTag(tag).
			SetArticleID(id)
	}

	err = a.db.ArticleTag.CreateBulk(bulk...).
		OnConflict().
		DoNothing().
		Exec(ctx)

	if err == nil {
		return nil
	}

	if errors.Is(err, sql.ErrNoRows) {
		log.Print(err)

		return nil
	}

	return errors.Wrap(err, "failed to save")
}

func (a Article) FindAll(ctx context.Context, limit int, offset int) ([]model.Article, error) {
	res, err := a.db.Article.Query().
		Where(
			article.DeletedAtIsNil(),
		).
		Order(ent.Asc(article.FieldCreatedAt)).
		Limit(limit).
		Offset(offset).
		All(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	articles := make([]model.Article, 0, len(res))

	for _, r := range res {
		articles = append(articles, model.Article{
			ID:          r.ID.String(),
			URL:         r.URL,
			Title:       r.Title,
			Thumbnail:   r.Thumbnail,
			Description: r.Description,
		})
	}

	// Tagを取得

	return articles, nil
}

func (a Article) LogicalDelete(ctx context.Context, id string) error {
	tempID, err := uuid.Parse(id)
	if err != nil {
		return errors.Wrap(err, "")
	}

	_, err = a.db.Article.UpdateOneID(tempID).
		SetDeletedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return errors.Wrap(err, "")
	}

	return nil
}
