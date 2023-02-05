//go:build e2e
// +build e2e

package auth_test

import (
	"context"
	"testing"

	"github.com/bufbuild/connect-go"
	"github.com/google/uuid"
	"github.com/morning-night-dream/platform/e2e/helper"
	authv1 "github.com/morning-night-dream/platform/pkg/connect/proto/auth/v1"
)

func TestE2EAuthChangePassword(t *testing.T) {
	t.Parallel()

	url := helper.GetEndpoint(t)

	t.Run("パスワード変更ができる", func(t *testing.T) {
		t.Parallel()

		newPassword := uuid.NewString()

		user := helper.NewUser(t, url)

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

	t.Run("元のパスワードが異なるためパスワード変更ができない", func(t *testing.T) {
		t.Parallel()

		newPassword := uuid.NewString()

		user := helper.NewUser(t, url)

		defer user.Delete(t)

		// パスワード変更
		req := &authv1.ChangePasswordRequest{
			Email:       user.EMail,
			OldPassword: "password",
			NewPassword: newPassword,
		}

		if _, err := user.Client.Auth.ChangePassword(context.Background(), connect.NewRequest(req)); err == nil {
			t.Errorf("success to change password: %s", err)
		}
	})

	t.Run("元のメアドが異なるためパスワード変更ができない", func(t *testing.T) {
		t.Parallel()

		newPassword := uuid.NewString()

		user := helper.NewUser(t, url)

		defer user.Delete(t)

		// パスワード変更
		req := &authv1.ChangePasswordRequest{
			Email:       "test@example.com",
			OldPassword: user.Password,
			NewPassword: newPassword,
		}

		if _, err := user.Client.Auth.ChangePassword(context.Background(), connect.NewRequest(req)); err == nil {
			t.Errorf("success to change password: %s", err)
		}
	})

	t.Run("Cookieがないためパスワード変更ができない", func(t *testing.T) {
		t.Parallel()

		newPassword := uuid.NewString()

		user := helper.NewUser(t, url)

		tmpClient := user.Client

		defer func() {
			user.Client = tmpClient
			user.Delete(t)
		}()

		user.Client = helper.NewPlainClient(t, url)

		// パスワード変更
		req := &authv1.ChangePasswordRequest{
			Email:       user.EMail,
			OldPassword: user.Password,
			NewPassword: newPassword,
		}

		if _, err := user.Client.Auth.ChangePassword(context.Background(), connect.NewRequest(req)); err == nil {
			t.Errorf("success to change password: %s", err)
		}
	})
}
