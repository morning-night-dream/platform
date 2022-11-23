package handler

import (
	"context"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/bufbuild/connect-go"
	"github.com/dyatlov/go-opengraph/opengraph"
	"github.com/google/uuid"
	"github.com/morning-night-dream/platform/app/core/database/store"
	"github.com/morning-night-dream/platform/app/core/model"
	"github.com/morning-night-dream/platform/app/core/proto"
	"github.com/morning-night-dream/platform/app/db/ent/proto/entpb"
	articlev1 "github.com/morning-night-dream/platform/pkg/api/article/v1"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Article struct {
	key    string
	client http.Client
	store  store.Article
	proto  *proto.Client
}

func NewArticle(
	store store.Article,
	proto *proto.Client,
) *Article {
	return &Article{
		key:    os.Getenv("API_KEY"),
		client: *http.DefaultClient,
		store:  store,
		proto:  proto,
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

	gr, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, ErrInvalidArgument
	}

	res, err := a.client.Do(gr.WithContext(ctx))
	if err != nil {
		return nil, ErrInternal
	}

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	og := opengraph.NewOpenGraph()

	err = og.ProcessHTML(strings.NewReader(string(body)))
	if err != nil {
		return nil, ErrInternal
	}

	thumbnail := ""
	if len(og.Images) > 0 {
		thumbnail = og.Images[0].URL
	}

	now := timestamppb.Now()

	if _, err := a.proto.Article.Create(
		ctx,
		&entpb.CreateArticleRequest{
			Article: &entpb.Article{
				Id:           []byte(uuid.New().String()),
				Title:        og.Title,
				Url:          og.URL,
				Description:  og.Description,
				Thumbnail:    thumbnail,
				CreatedAt:    now,
				UpdatedAt:    now,
				Tags:         []*entpb.ArticleTag{},
				ReadArticles: []*entpb.ReadArticle{},
			},
		},
	); err != nil {
		log.Print(err)

		return nil, ErrInternal
	}

	return connect.NewResponse(&articlev1.ShareResponse{}), nil
}

func (a *Article) List(
	ctx context.Context,
	req *connect.Request[articlev1.ListRequest],
) (*connect.Response[articlev1.ListResponse], error) {
	limit := int(req.Msg.MaxPageSize)

	items, err := a.proto.Article.List(ctx, &entpb.ListArticleRequest{
		PageSize:  int32(limit),
		PageToken: req.Msg.PageToken,
		View:      1,
	})
	if err != nil {
		log.Print(err)
		return nil, errors.Wrap(err, "")
	}

	articles := make([]*articlev1.Article, 0, len(items.ArticleList))

	for _, item := range items.ArticleList {
		tags := make([]string, len(item.Tags))
		for i, tag := range item.Tags {
			tags[i] = tag.Tag
		}

		articles = append(articles, &articlev1.Article{
			Id:          string(item.Id),
			Title:       item.Title,
			Url:         item.Url,
			Description: item.Description,
			Thumbnail:   item.Thumbnail,
			Tags:        tags,
		})
	}

	res := connect.NewResponse(&articlev1.ListResponse{
		Articles:      articles,
		NextPageToken: items.NextPageToken,
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
