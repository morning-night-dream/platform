package handler

import (
	"context"
	"encoding/base64"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/bufbuild/connect-go"
	"github.com/google/uuid"
	"github.com/morning-night-dream/platform/internal/database/store"
	"github.com/morning-night-dream/platform/internal/model"
	articlev1 "github.com/morning-night-dream/platform/pkg/proto/article/v1"
	"github.com/pkg/errors"
)

type Article struct {
	key    string
	client http.Client
	store  store.Article
}

func NewArticle(
	store store.Article,
) *Article {
	return &Article{
		key:    os.Getenv("API_KEY"),
		client: *http.DefaultClient,
		store:  store,
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

	if err := a.store.Save(ctx, model.Article{
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

	items, err := a.store.FindAll(ctx, limit, offset)
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
	if err := a.store.LogicalDelete(ctx, req.Msg.Id); err != nil {
		return nil, errors.Wrap(err, "")
	}

	return connect.NewResponse(&articlev1.DeleteResponse{}), nil
}

func (a *Article) Read(
	ctx context.Context,
	req *connect.Request[articlev1.ReadRequest],
) (*connect.Response[articlev1.ReadResponse], error) {
	jwt := req.Header().Get("Authorization")

	ctx, err := model.Authorize(ctx, jwt)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	if err := a.store.SaveRead(ctx, req.Msg.Id, model.GetUIDCtx(ctx)); err != nil {
		return nil, errors.Wrap(err, "")
	}

	return connect.NewResponse(&articlev1.ReadResponse{}), nil
}

func (a *Article) AddTag(
	ctx context.Context,
	req *connect.Request[articlev1.AddTagRequest],
) (*connect.Response[articlev1.AddTagResponse], error) {
	item, err := a.store.Find(ctx, req.Msg.Id)
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

	if err := a.store.Save(ctx, item); err != nil {
		return nil, errors.Wrap(err, "")
	}

	item.Tags = tags

	return connect.NewResponse(&articlev1.AddTagResponse{}), nil
}

func (a *Article) ListTag(
	ctx context.Context,
	req *connect.Request[articlev1.ListTagRequest],
) (*connect.Response[articlev1.ListTagResponse], error) {
	tags, err := a.store.FindAllTag(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	return connect.NewResponse(&articlev1.ListTagResponse{
		Tags: tags,
	}), nil
}
