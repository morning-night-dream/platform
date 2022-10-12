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
	"github.com/morning-night-dream/article-share/app/core/database/store"
	"github.com/morning-night-dream/article-share/app/core/model"
	articlev1 "github.com/morning-night-dream/article-share/pkg/api/article/v1"
	"github.com/pkg/errors"
)

type ArticleHandler struct {
	key    string
	client http.Client
	store  store.Article
}

func NewArticleHandler(store store.Article) *ArticleHandler {
	return &ArticleHandler{
		key:    os.Getenv("API_KEY"),
		client: *http.DefaultClient,
		store:  store,
	}
}

func (s *ArticleHandler) Share(
	ctx context.Context,
	req *connect.Request[articlev1.ShareRequest],
) (*connect.Response[articlev1.ShareResponse], error) {
	if req.Header().Get("X-API-KEY") != s.key {
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

	res, err := s.client.Do(gr.WithContext(ctx))
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

	imageURL := ""
	if len(og.Images) > 0 {
		imageURL = og.Images[0].URL
	}

	article := model.Article{
		URL:         og.URL,
		Title:       og.Title,
		ImageURL:    imageURL,
		Description: og.Description,
	}

	err = s.store.Save(ctx, article)

	if err != nil {
		log.Print(err)

		return nil, ErrInternal
	}

	return connect.NewResponse(&articlev1.ShareResponse{}), nil
}

func (s *ArticleHandler) List(
	ctx context.Context,
	req *connect.Request[articlev1.ListRequest],
) (*connect.Response[articlev1.ListResponse], error) {
	limit := 100

	items, err := s.store.FindAll(ctx, limit, int(req.Msg.Page)*limit)
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
			ImageUrl:    item.ImageURL,
		})
	}

	res := connect.NewResponse(&articlev1.ListResponse{
		Articles: articles,
	})

	return res, nil
}
