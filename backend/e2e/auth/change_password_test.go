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
	authv1 "github.com/morning-night-dream/platform/pkg/proto/auth/v1"
)

func TestE2EAuthChangePassword(t *testing.T) {
	t.Parallel()

	url := helper.GetEndpoint(t)

	t.Run("パスワード変更ができる", func(t *testing.T) {
		t.Parallel()

		email := fmt.Sprintf("%s@example.com", uuid.NewString())

		password := uuid.NewString()

		newPassword := uuid.NewString()

		user := helper.NewUser(t, email, password, url)

		defer func() {
			user.Password = newPassword
			user.Delete(t)
		}()

		// パスワード変更
		req := &authv1.ChangePasswordRequest{
			Email:       user.EMail,
			OldPassword: user.Password,
			NewPassword: newPassword,
		}

		if _, err := user.Client.Auth.ChangePassword(context.Background(), connect.NewRequest(req)); err != nil {
			t.Errorf("failed to change password: %s", err)
		}

		// 元のパスワードでサインインできない
		client := helper.NewPlainClient(t, url)

		if _, err := client.Auth.SignIn(context.Background(), connect.NewRequest(&authv1.SignInRequest{
			Email:    user.EMail,
			Password: user.Password,
		})); err == nil {
			t.Error("success to sign in")
		}
	})
}
