//go:build e2e
// +build e2e

package auth_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/bufbuild/connect-go"
	"github.com/google/uuid"
	"github.com/morning-night-dream/platform/e2e/helper"
	authv1 "github.com/morning-night-dream/platform/pkg/connect/proto/auth/v1"
)

func TestE2EAuthDelete(t *testing.T) {
	t.Parallel()

	url := helper.GetCoreEndpoint(t)

	t.Run("ユーザーを削除できる", func(t *testing.T) {
		t.Parallel()

		client := helper.NewPlainConnectClient(t, url)

		id := uuid.New().String()
		email := fmt.Sprintf("%s@example.com", id)
		password := id

		sureq := &authv1.SignUpRequest{
			Email:    email,
			Password: password,
		}
		if _, err := client.Auth.SignUp(context.Background(), connect.NewRequest(sureq)); err != nil {
			t.Fatalf("failed to auth sign up: %s", err)
		}

		sireq := &authv1.SignInRequest{
			Email:    email,
			Password: password,
		}

		res, err := client.Auth.SignIn(context.Background(), connect.NewRequest(sireq))
		if err != nil {
			t.Fatalf("failed to auth sign in: %s", err)
		}

		cookie := res.Header().Get("Set-Cookie")

		dclient := helper.NewConnectClientWithCookie(t, cookie, url)

		req := &authv1.DeleteRequest{
			Email:    email,
			Password: password,
		}

		if _, err := dclient.Auth.Delete(context.Background(), connect.NewRequest(req)); err != nil {
			t.Fatalf("failed to delete user in: %s", err)
		}
	})
}
