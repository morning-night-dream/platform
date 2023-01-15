//go:build e2e
// +build e2e

package article_test

import (
	"context"
	"log"
	"net/http"
	"testing"

	"github.com/bufbuild/connect-go"
	"github.com/morning-night-dream/platform/e2e/helper"
	authv1 "github.com/morning-night-dream/platform/pkg/proto/auth/v1"
)

func TestE2EAuthSignOut(t *testing.T) {
	t.Parallel()

	url := helper.GetEndpoint(t)

	t.Run("サインアウトできる", func(t *testing.T) {
		t.Parallel()

		hc := helper.NewClient(t, http.DefaultClient, url)

		sreq := &authv1.SignInRequest{
			Email:    helper.GetEMail(t),
			Password: helper.GetPassword(t),
		}

		sres, err := hc.Auth.SignIn(context.Background(), connect.NewRequest(sreq))
		if err != nil {
			t.Fatalf("failed to auth sign in: %s", err)
		}

		c := &http.Client{
			Transport: helper.NewCookieTransport(t, sres.Header().Get("Set-Cookie")),
		}

		client := helper.NewClient(t, c, url)

		req := &authv1.SignOutRequest{}

		res, err := client.Auth.SignOut(context.Background(), connect.NewRequest(req))
		if err != nil {
			t.Fatalf("failed to auth sign out: %s", err)
		}

		log.Println(res.Header().Get("Set-Cookie"))
	})
}
