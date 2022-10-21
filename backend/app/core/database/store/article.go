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
	"github.com/morning-night-dream/platform/pkg/ent/articletag"
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

func (a Article) Save(ctx context.Context, item model.Article) error {
	id, err := uuid.Parse(item.ID)
	if err != nil {
		return errors.Wrap(err, "failed to parse uuid")
	}

	tx, err := a.db.Tx(ctx)
	if err != nil {
		return errors.Wrap(err, "starting a transaction")
	}

	now := time.Now().UTC()

	err = tx.Article.Create().
		SetID(id).
		SetTitle(item.Title).
		SetDescription(item.Description).
		SetURL(item.URL).
		SetThumbnail(item.Thumbnail).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		OnConflictColumns(article.FieldURL).
		UpdateTitle().
		UpdateDescription().
		UpdateThumbnail().
		UpdateUpdatedAt().
		Exec(ctx)
	if err != nil {
		log.Printf("failed to save article %s", err)

		if re := tx.Rollback(); re != nil {
			return errors.Wrap(re, "")
		}

		return errors.Wrap(err, "")
	}

	if len(item.Tags) == 0 {
		if ce := tx.Commit(); ce != nil {
			return errors.Wrap(ce, "")
		}

		return nil
	}

	if _, err = tx.ArticleTag.Delete().Where(articletag.ArticleIDEQ(id)).Exec(ctx); err != nil {
		if re := tx.Rollback(); re != nil {
			return errors.Wrap(re, "")
		}

		return errors.Wrap(err, "")
	}

	bulk := make([]*ent.ArticleTagCreate, len(item.Tags))
	for i, tag := range item.Tags {
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
		if ce := tx.Commit(); ce != nil {
			return errors.Wrap(ce, "")
		}

		return nil
	}

	if errors.Is(err, sql.ErrNoRows) {
		log.Print(err)

		if re := tx.Rollback(); re != nil {
			return errors.Wrap(re, "")
		}

		return errors.Wrap(err, "")
	}

	log.Printf("failed to save article tags %s", err)

	if re := tx.Rollback(); re != nil {
		return errors.Wrap(re, "")
	}

	return errors.Wrap(err, "")
}

func (a Article) Find(ctx context.Context, id string) (model.Article, error) {
	tid, err := uuid.Parse(id)
	if err != nil {
		return model.Article{}, errors.Wrap(err, "failed to parse uuid")
	}

	item, err := a.db.Article.Query().
		WithTags().
		Where(article.IDEQ(tid)).
		First(ctx)
	if err != nil {
		return model.Article{}, errors.Wrap(err, "failed to find article")
	}

	tags := make([]string, len(item.Edges.Tags))

	for i, tag := range item.Edges.Tags {
		tags[i] = tag.Tag
	}

	return model.Article{
		ID:          item.ID.String(),
		URL:         item.URL,
		Title:       item.Title,
		Thumbnail:   item.Thumbnail,
		Description: item.Description,
		Tags:        tags,
	}, nil
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

	return articles, nil
}

func (a Article) FindAllTag(ctx context.Context) ([]string, error) {
	tags, err := a.db.ArticleTag.
		Query().
		Unique(true).
		Select(articletag.FieldTag).
		Strings(ctx)
	if err != nil {
		return []string{}, errors.Wrap(err, "")
	}

	return tags, nil
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
