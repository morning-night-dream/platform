//go:build e2e
// +build e2e

package article_test

import (
	"context"
	"encoding/json"
	"io"
	"reflect"
	"testing"

	"github.com/morning-night-dream/platform/e2e/helper"
	"github.com/morning-night-dream/platform/pkg/openapi"
)

func TestGatewayE2EArticleList(t *testing.T) {
	t.Parallel()

	size := uint32(10)

	helper.BulkInsert(t, int(size))

	url := helper.GetGatewayEndpoint(t)

	t.Run("記事が一覧できる", func(t *testing.T) {
		t.Parallel()

		client := helper.NewOpenAPIClient(t, url)

		res, err := client.Client.V1ListArticles(context.Background(), &openapi.V1ListArticlesParams{
			PageToken:   nil,
			MaxPageSize: 20,
		})
		if err != nil {
			t.Fatalf("failed to list article: %s", err)
		}

		defer res.Body.Close()

		body, _ := io.ReadAll(res.Body)

		var article openapi.ListArticleResponse
		if err := json.Unmarshal(body, &article); err != nil {
			t.Fatalf("failed marshal response: %s caused by %s", body, err)
			return
		}

		if !reflect.DeepEqual(len(*article.Articles), int(size)) {
			t.Errorf("Articles length = %v, want %v", len(*article.Articles), size)
		}
	})
}
