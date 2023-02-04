package controller

import (
	"encoding/json"
	"net/http"

	"github.com/bufbuild/connect-go"
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/google/uuid"
	"github.com/morning-night-dream/platform/pkg/openapi"
	articlev1 "github.com/morning-night-dream/platform/pkg/proto/article/v1"
)

func (c Controller) V1ListArticles(w http.ResponseWriter, r *http.Request, params openapi.V1ListArticlesParams) {
	pageToken := ""
	if params.PageToken != nil {
		pageToken = *params.PageToken
	}

	req := &articlev1.ListRequest{
		PageToken:   pageToken,
		MaxPageSize: uint32(params.MaxPageSize),
	}
	res, err := c.client.Article.List(r.Context(), connect.NewRequest(req))
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte(err.Error()))

		return
	}

	articles := make([]openapi.Article, len(res.Msg.Articles))

	for i, item := range res.Msg.Articles {
		uid, _ := uuid.Parse(item.Id)

		id := openapi_types.UUID(uid)

		articles[i] = openapi.Article{
			Id:          &id,
			Title:       &item.Title,
			Url:         &item.Url,
			Description: &item.Description,
			Thumbnail:   &item.Thumbnail,
			Tags:        &item.Tags,
		}
	}

	rs := openapi.ListArticleResponse{
		Articles:      &articles,
		NextPageToken: &res.Msg.NextPageToken,
	}

	if err := json.NewEncoder(w).Encode(rs); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
