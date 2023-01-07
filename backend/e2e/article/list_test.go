//go:build e2e
// +build e2e

package article_test

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/bufbuild/connect-go"
	"github.com/morning-night-dream/platform/e2e/article"
	"github.com/morning-night-dream/platform/e2e/helper"
	articlev1 "github.com/morning-night-dream/platform/pkg/proto/article/v1"
)

func TestE2EArticleList(t *testing.T) {
	t.Parallel()

	size := uint32(10)

	helper.BulkInsert(t, int(size))

	url := helper.GetEndpoint(t)

	t.Run("記事が一覧できる", func(t *testing.T) {
		t.Parallel()

		hc := &http.Client{
			Transport: helper.NewAPIKeyTransport(t, helper.GetAPIKey(t)),
		}

		client := article.New(t, hc, url)

		req := &articlev1.ListRequest{
			MaxPageSize: size,
		}

		res, err := client.Article.List(context.Background(), connect.NewRequest(req))
		if err != nil {
			t.Fatalf("failed to article share: %s", err)
		}

		if !reflect.DeepEqual(len(res.Msg.Articles), int(size)) {
			t.Errorf("Articles length = %v, want %v", len(res.Msg.Articles), size)
		}
	})
}
