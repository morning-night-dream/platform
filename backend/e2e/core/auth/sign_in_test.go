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
	authv1 "github.com/morning-night-dream/platform/pkg/connect/proto/auth/v1"
)

func TestE2EAuthSignIn(t *testing.T) {
	t.Parallel()

	url := helper.GetCoreEndpoint(t)

	t.Run("サインインできる", func(t *testing.T) {
		t.Parallel()

		user := helper.NewUser(t, url)

		defer func() {
			user.Delete(t)
		}()

		client := helper.NewConnectClient(t, http.DefaultClient, url)

		req := &authv1.SignInRequest{
			Email:    user.EMail,
			Password: user.Password,
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
