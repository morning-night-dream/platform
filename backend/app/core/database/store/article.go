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

	tx, err := a.db.Tx(ctx)
	if err != nil {
		return errors.Wrap(err, "starting a transaction")
	}

	now := time.Now().UTC()

	err = tx.Article.Create().
		SetID(id).
		SetTitle(article.Title).
		SetDescription(article.Description).
		SetURL(article.URL).
		SetThumbnail(article.Thumbnail).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		OnConflict().
		DoNothing().
		Exec(ctx)
	if err != nil {
		// https://github.com/ent/ent/issues/2176 により、
		// on conflict do nothingとしてもerror no rowsが返るため、個別にハンドリングする
		if errors.Is(err, sql.ErrNoRows) {
			log.Print(err)

			return tx.Commit()
		}

		log.Printf("failed to save article %s", err)

		return tx.Rollback()
	}

	article.Tags = []string{"Go"}

	if len(article.Tags) == 0 {
		return tx.Commit()
	}

	bulk := make([]*ent.ArticleTagCreate, len(article.Tags))
	for i, tag := range article.Tags {
		bulk[i] = tx.ArticleTag.Create().
			SetTag(tag).
			SetArticleID(id).
			SetCreatedAt(now).
			SetUpdatedAt(now)
	}

	err = tx.ArticleTag.CreateBulk(bulk...).
		OnConflict().
		DoNothing().
		Exec(ctx)

	if err == nil {
		return tx.Commit()
	}

	if errors.Is(err, sql.ErrNoRows) {
		log.Print(err)

		return tx.Rollback()
	}

	log.Printf("failed to save article tags %s", err)

	return tx.Rollback()
}

func (a Article) FindAll(ctx context.Context, limit int, offset int) ([]model.Article, error) {
	res, err := a.db.Article.Query().
		WithTags().
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
		tags := make([]string, 0, len(r.Edges.Tags))
		for _, t := range r.Edges.Tags {
			tags = append(tags, t.Tag)
		}

		articles = append(articles, model.Article{
			ID:          r.ID.String(),
			URL:         r.URL,
			Title:       r.Title,
			Thumbnail:   r.Thumbnail,
			Description: r.Description,
			Tags:        tags,
		})
	}

	// Tagを取得

	return articles, nil
}

func (a Article) LogicalDelete(ctx context.Context, id string) error {
	tid, err := uuid.Parse(id)
	if err != nil {
		return errors.Wrap(err, "")
	}

	_, err = a.db.Article.UpdateOneID(tid).
		SetDeletedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return errors.Wrap(err, "")
	}

	return nil
}

func (a Article) SaveRead(ctx context.Context, id, uid string) error {
	tid, err := uuid.Parse(id)
	if err != nil {
		return errors.Wrap(err, "")
	}

	tuid, err := uuid.Parse(uid)
	if err != nil {
		return errors.Wrap(err, "")
	}

	now := time.Now().UTC()

	err = a.db.ReadArticle.Create().
		SetID(uuid.New()).
		SetUserID(tuid).
		SetArticleID(tid).
		SetReadAt(now).
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

		log.Printf("failed to save article %s", err)

		return errors.Wrap(err, "failed to save read article")
	}

	return nil
}
