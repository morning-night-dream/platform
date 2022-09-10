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
	articlev1 "github.com/morning-night-dream/article-share/api/article/v1"
	"github.com/morning-night-dream/article-share/model"
)

type ArticleHandler struct {
	key    string
	client http.Client
	store  model.ArticleStore
}

func NewArticleHandler(store model.ArticleStore) *ArticleHandler {
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