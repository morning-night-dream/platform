package handler

import (
	"context"
	"encoding/base64"
	"log"
	"net/url"
	"os"
	"strconv"

	"github.com/bufbuild/connect-go"
	"github.com/google/uuid"
	"github.com/morning-night-dream/platform/internal/adapter/gateway"
	"github.com/morning-night-dream/platform/internal/domain/model"
	articlev1 "github.com/morning-night-dream/platform/pkg/proto/article/v1"
	"github.com/pkg/errors"
)

type Article struct {
	key     string
	article *gateway.Article
	handle  *Handle
}

func NewArticle(
	article *gateway.Article,
	handle *Handle,
) *Article {
	return &Article{
		key:     os.Getenv("API_KEY"),
		article: article,
		handle:  handle,
	}
}

func (a *Article) Share(
	ctx context.Context,
	req *connect.Request[articlev1.ShareRequest],
) (*connect.Response[articlev1.ShareResponse], error) {
	if req.Header().Get("X-API-KEY") != a.key {
		return nil, ErrUnauthorized
	}

	u, err := url.Parse(req.Msg.Url)
	if err != nil {
		return nil, ErrInvalidArgument
	}

	id := uuid.NewString()

	if err := a.article.Save(ctx, model.Article{
		ID:          id,
		URL:         u.String(),
		Title:       req.Msg.Title,
		Thumbnail:   req.Msg.Thumbnail,
		Description: req.Msg.Description,
	}); err != nil {
		log.Print(err)

		return nil, ErrInternal
	}

	res := &articlev1.ShareResponse{
		Article: &articlev1.Article{
			Id:          id,
			Url:         u.String(),
			Title:       req.Msg.Title,
			Thumbnail:   req.Msg.Thumbnail,
			Description: req.Msg.Description,
		},
	}

	return connect.NewResponse(res), nil
}

func (a *Article) List(
	ctx context.Context,
	req *connect.Request[articlev1.ListRequest],
) (*connect.Response[articlev1.ListResponse], error) {
	limit := int(req.Msg.MaxPageSize)

	dec, err := base64.StdEncoding.DecodeString(req.Msg.PageToken)
	if err != nil {
		dec = []byte("0")
	}

	offset, err := strconv.Atoi(string(dec))
	if err != nil {
		offset = 0
	}

	items, err := a.article.FindAll(ctx, limit, offset)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	articles := make([]*articlev1.Article, 0, len(items))

	for _, item := range items {
		articles = append(articles, &articlev1.Article{
			Id:          item.ID,
			Title:       item.Title,
			Url:         item.URL,
			Description: item.Description,
			Thumbnail:   item.Thumbnail,
			Tags:        item.Tags,
		})
	}

	token := base64.StdEncoding.EncodeToString([]byte(strconv.Itoa(offset + limit)))
	if len(articles) < limit {
		token = ""
	}

	res := connect.NewResponse(&articlev1.ListResponse{
		Articles:      articles,
		NextPageToken: token,
	})

	return res, nil
}

func (a *Article) Delete(
	ctx context.Context,
	req *connect.Request[articlev1.DeleteRequest],
) (*connect.Response[articlev1.DeleteResponse], error) {
	if err := a.article.LogicalDelete(ctx, req.Msg.Id); err != nil {
		return nil, errors.Wrap(err, "")
	}

	return connect.NewResponse(&articlev1.DeleteResponse{}), nil
}

func (a *Article) Read(
	ctx context.Context,
	req *connect.Request[articlev1.ReadRequest],
) (*connect.Response[articlev1.ReadResponse], error) {
	auth, err := a.handle.Authorize(ctx, req.Header())
	if err != nil {
		return nil, ErrUnauthorized
	}

	if err := a.article.SaveRead(ctx, req.Msg.Id, auth.UserID); err != nil {
		return nil, errors.Wrap(err, "")
	}

	return connect.NewResponse(&articlev1.ReadResponse{}), nil
}

func (a *Article) AddTag(
	ctx context.Context,
	req *connect.Request[articlev1.AddTagRequest],
) (*connect.Response[articlev1.AddTagResponse], error) {
	item, err := a.article.Find(ctx, req.Msg.Id)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	item.Tags = append(item.Tags, req.Msg.Tag)

	tmp := make(map[string]struct{})

	for _, tag := range item.Tags {
		tmp[tag] = struct{}{}
	}

	tags := make([]string, 0, len(tmp))
	for i := range tmp {
		tags = append(tags, i)
	}

	if err := a.article.Save(ctx, item); err != nil {
		return nil, errors.Wrap(err, "")
	}

	item.Tags = tags

	return connect.NewResponse(&articlev1.AddTagResponse{}), nil
}

func (a *Article) ListTag(
	ctx context.Context,
	req *connect.Request[articlev1.ListTagRequest],
) (*connect.Response[articlev1.ListTagResponse], error) {
	tags, err := a.article.FindAllTag(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	return connect.NewResponse(&articlev1.ListTagResponse{
		Tags: tags,
	}), nil
}
