//go:build e2e
// +build e2e

package article_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/bufbuild/connect-go"
	"github.com/morning-night-dream/platform/e2e/helper"
	articlev1 "github.com/morning-night-dream/platform/pkg/proto/article/v1"
	authv1 "github.com/morning-night-dream/platform/pkg/proto/auth/v1"
)

func TestE2EArticleRead(t *testing.T) {
	t.Parallel()

	size := uint32(10)

	helper.BulkInsert(t, int(size))

	url := helper.GetEndpoint(t)

	t.Run("記事が既読できる", func(t *testing.T) {
		t.Parallel()

		ac := helper.NewClient(t, http.DefaultClient, url)

		sreq := &authv1.SignInRequest{
			Email:    helper.GetEMail(t),
			Password: helper.GetPassword(t),
		}

		sres, _ := ac.Auth.SignIn(context.Background(), connect.NewRequest(sreq))

		hc := &http.Client{
			Transport: helper.NewCookieTransport(t, sres.Header().Get("Set-Cookie")),
		}

		client := helper.NewClient(t, hc, url)

		articles, err := client.Article.List(context.Background(), connect.NewRequest(&articlev1.ListRequest{
			MaxPageSize: size,
		}))
		if err != nil {
			t.Fatalf("failed to article share: %s", err)
		}

		req := &articlev1.ReadRequest{
			Id: articles.Msg.Articles[0].Id,
		}

		_, err = client.Article.Read(context.Background(), connect.NewRequest(req))
		if err != nil {
			t.Fatalf("failed to article share: %s", err)
		}
	})
}
