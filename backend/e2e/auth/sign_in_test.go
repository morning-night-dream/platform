//go:build e2e
// +build e2e

package auth_test

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/bufbuild/connect-go"
	"github.com/morning-night-dream/platform/e2e/helper"
	authv1 "github.com/morning-night-dream/platform/pkg/proto/auth/v1"
)

func TestE2EAuthSignIn(t *testing.T) {
	t.Parallel()

	url := helper.GetEndpoint(t)

	t.Run("サインインできる", func(t *testing.T) {
		t.Parallel()

		client := helper.NewClient(t, http.DefaultClient, url)

		req := &authv1.SignInRequest{
			Email:    helper.GetEMail(t),
			Password: helper.GetPassword(t),
		}

		res, err := client.Auth.SignIn(context.Background(), connect.NewRequest(req))
		if err != nil {
			t.Fatalf("failed to auth sign in: %s", err)
		}

		if reflect.DeepEqual(len(res.Header().Get("Set-Cookie")), 0) {
			t.Errorf("cookie is empty")
		}
	})
}
