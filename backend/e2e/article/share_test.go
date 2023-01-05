//go:build e2e
// +build e2e

package article_test

import (
	"context"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/bufbuild/connect-go"
	"github.com/morning-night-dream/platform/e2e/article"
	"github.com/morning-night-dream/platform/e2e/helper"
	articlev1 "github.com/morning-night-dream/platform/pkg/proto/article/v1"
)

func TestE2EArticleShare(t *testing.T) {
	t.Parallel()

	url := helper.GetEndpoint(t)

	t.Run("記事が共有できる", func(t *testing.T) {
		t.Parallel()

		hc := &http.Client{
			Transport: helper.NewAPIKeyTransport(t, "local"),
		}

		client := article.New(t, hc, url)

		req := &articlev1.ShareRequest{
			Url:         "http://www.example.com",
			Title:       "title",
			Description: "description",
			Thumbnail:   "http://www.example.com/thumbnail.jpg",
		}

		res, err := client.Article.Share(context.Background(), connect.NewRequest(req))
		if err != nil {
			t.Fatalf("faile to article share: %s", err)
		}

		if !reflect.DeepEqual(res.Msg.Article.Url, req.Url) {
			t.Errorf("Url = %v, want %v", res.Msg.Article.Url, req.Url)
		}
		if !reflect.DeepEqual(res.Msg.Article.Title, req.Title) {
			t.Errorf("Title = %v, want %v", res.Msg.Article.Title, req.Title)
		}
		if !reflect.DeepEqual(res.Msg.Article.Description, req.Description) {
			t.Errorf("Description = %v, want %v", res.Msg.Article.Description, req.Description)
		}
		if !reflect.DeepEqual(res.Msg.Article.Thumbnail, req.Thumbnail) {
			t.Errorf("Thumbnail = %v, want %v", res.Msg.Article.Thumbnail, req.Thumbnail)
		}
	})

	t.Run("Api-Keyがなくて記事が共有できない", func(t *testing.T) {
		t.Parallel()

		hc := &http.Client{}

		client := article.New(t, hc, url)

		req := &articlev1.ShareRequest{
			Url:         "http://www.example.com",
			Title:       "title",
			Description: "description",
			Thumbnail:   "http://www.example.com/thumbnail.jpg",
		}

		_, err := client.Article.Share(context.Background(), connect.NewRequest(req))
		if !strings.Contains(err.Error(), "Unauthenticated") {
			t.Errorf("err = %v", err)
		}
		if !strings.Contains(err.Error(), "unauthorized") {
			t.Errorf("err = %v", err)
		}
	})
}
