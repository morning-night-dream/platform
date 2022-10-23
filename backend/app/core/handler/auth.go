package handler

import (
	"context"
	"log"

	"github.com/bufbuild/connect-go"
	"github.com/google/uuid"
	"github.com/morning-night-dream/platform/app/core/model"
	"github.com/morning-night-dream/platform/app/core/store"
	authv1 "github.com/morning-night-dream/platform/pkg/api/auth/v1"
)

type Auth struct {
	store store.Auth
}

func NewAuth(store store.Auth) *Auth {
	return &Auth{
		store: store,
	}
}

func (a Auth) SignUp(
	ctx context.Context,
	req *connect.Request[authv1.SignUpRequest],
) (*connect.Response[authv1.SignUpResponse], error) {
	email := req.Msg.Email
	if email == "" {
		log.Printf("fail to sign up caused by invalid email %s", email)

		return nil, ErrInvalidArgument
	}

	auth := model.Auth{
		UserID:   uuid.NewString(),
		Email:    email,
		Password: "password",
	}

	err := a.store.Save(ctx, auth)
	if err != nil {
		log.Printf("fail to sign up caused by %v", err)

		return nil, ErrInternal
	}

	return connect.NewResponse(&authv1.SignUpResponse{}), nil
}

func (a Auth) SignIn(
	ctx context.Context,
	req *connect.Request[authv1.SignInRequest],
) (*connect.Response[authv1.SignInResponse], error) {
	email := req.Msg.Email
	if email == "" {
		return nil, ErrUnauthorized
	}

	password := req.Msg.Password
	if password == "" {
		return nil, ErrUnauthorized
	}

	tokens, err := a.store.FindFromEmailPass(ctx, email, password)
	if err != nil {
		log.Printf("fail to sign in caused by %v", err)

		return nil, ErrUnauthorized
	}

	res := &authv1.SignInResponse{
		IdToken:      tokens.IDToken,
		RefreshToken: tokens.RefreshToken,
	}

	return connect.NewResponse(res), nil
}

func (a Auth) Refresh(
	ctx context.Context,
	req *connect.Request[authv1.RefreshRequest],
) (*connect.Response[authv1.RefreshResponse], error) {

	tokenPair, err := a.store.FindFromIDToken(ctx, req.Msg.RefreshToken)
	if err != nil {
		return nil, ErrUnauthorized
	}

	res := &authv1.RefreshResponse{
		IdToken:      tokenPair.IDToken,
		RefreshToken: tokenPair.RefreshToken,
	}

	return connect.NewResponse(res), nil
}

func (a Auth) ChangePassword(
	ctx context.Context,
	req *connect.Request[authv1.ChangePasswordRequest],
) (*connect.Response[authv1.ChangePasswordResponse], error) {
	return nil, nil
}
