package handler

import (
	"context"
	"encoding/base64"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/bufbuild/connect-go"
	"github.com/dyatlov/go-opengraph/opengraph"
	"github.com/morning-night-dream/platform/app/core/database/store"
	"github.com/morning-night-dream/platform/app/core/model"
	articlev1 "github.com/morning-night-dream/platform/pkg/api/article/v1"
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

func (a *ArticleHandler) Share(
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

	if err := a.store.Save(ctx, article); err != nil {
		log.Print(err)

		return nil, ErrInternal
	}

	return connect.NewResponse(&articlev1.ShareResponse{}), nil
}

func (a *ArticleHandler) List(
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
			ImageUrl:    item.ImageURL,
		})
	}

	token := base64.StdEncoding.EncodeToString([]byte(strconv.Itoa(offset*limit + 1)))
	if len(articles) < limit {
		token = ""
	}

	res := connect.NewResponse(&articlev1.ListResponse{
		Articles:      articles,
		NextPageToken: token,
	})

	return res, nil
}

func (a *ArticleHandler) Delete(
	ctx context.Context,
	req *connect.Request[articlev1.DeleteRequest],
) (*connect.Response[articlev1.DeleteResponse], error) {
	if err := a.store.LogicalDelete(ctx, req.Msg.Id); err != nil {
		return nil, errors.Wrap(err, "")
	}

	return connect.NewResponse(&articlev1.DeleteResponse{}), nil
}
