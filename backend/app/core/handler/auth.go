package handler

import (
	"context"
	"log"

	"github.com/bufbuild/connect-go"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/morning-night-dream/platform/app/core/database/store"
	"github.com/morning-night-dream/platform/app/core/model"
	authv1 "github.com/morning-night-dream/platform/pkg/api/auth/v1"
)

type AuthHandler struct {
	store store.Auth
}

func NewAuthHandler(store store.Auth) *AuthHandler {
	return &AuthHandler{
		store: store,
	}
}

func (a AuthHandler) SignUp(
	ctx context.Context,
	req *connect.Request[authv1.SignUpRequest],
) (*connect.Response[authv1.SignUpResponse], error) {
	email := req.Msg.Email
	if email == "" {
		log.Printf("fail to sign up caused by invalid email %s", email)

		return nil, ErrInvalidArgument
	}

	password := req.Msg.Password
	if password == "" {
		log.Printf("fail to sign up caused by invalid password %s", password)

		return nil, ErrInvalidArgument
	}

	loginID := req.Msg.LoginId
	if loginID == "" {
		log.Printf("fail to sign up caused by invalid loginId %s", loginID)

		return nil, ErrInvalidArgument
	}

	auth := model.Auth{
		UserID:   uuid.NewString(),
		Email:    email,
		Password: password,
		LoginID:  loginID,
	}

	err := a.store.Save(ctx, auth)
	if err != nil {
		log.Printf("fail to sign up caused by %v", err)

		return nil, ErrInternal
	}

	log.Printf("success to sign up")

	return connect.NewResponse(&authv1.SignUpResponse{}), nil
}

func (a AuthHandler) SignIn(
	ctx context.Context,
	req *connect.Request[authv1.SignInRequest],
) (*connect.Response[authv1.SignInResponse], error) {
	loginID := req.Msg.LoginId
	if loginID == "" {
		return nil, ErrUnauthorized
	}

	password := req.Msg.Password
	if password == "" {
		return nil, ErrUnauthorized
	}

	auth, err := a.store.FindFromIDPass(ctx, loginID, password)
	if err != nil {
		log.Printf("fail to sign in caused by %v", err)

		return nil, ErrUnauthorized
	}

	tokenPair, err := model.GenerateToken(auth.UserID)
	if err != nil {
		log.Printf("fail to sign in caused by %v", err)

		return nil, ErrUnauthorized
	}

	res := &authv1.SignInResponse{
		IdToken:      tokenPair.IDToken.String(),
		RefreshToken: tokenPair.RefreshToken.String(),
	}

	log.Printf("success to sign in")

	return connect.NewResponse(res), nil
}

func (a AuthHandler) Refresh(
	ctx context.Context,
	req *connect.Request[authv1.RefreshRequest],
) (*connect.Response[authv1.RefreshResponse], error) {
	idTokenString := req.Header().Get("Authorization")

	ctx, err := model.Authorize(ctx, idTokenString)
	if err != nil {
		log.Printf("fail to refresh caused by %v", err)

		return nil, ErrUnauthorized
	}

	uid := model.GetUIDCtx(ctx)

	refreshToken, err := model.CreateTokenFrom(req.Msg.RefreshToken)
	if err != nil {
		log.Printf("fail to refresh caused by %v", err)

		return nil, ErrUnauthorized
	}

	claims, ok := refreshToken.Claims.(jwt.MapClaims)
	if ok && !refreshToken.Valid {
		log.Printf("fail to refresh caused by invalid refresh token")

		return nil, ErrUnauthorized
	}

	if sub, _ := claims["sub"].(string); sub != uid {
		log.Printf("fail to refresh caused by invalid uid")

		return nil, ErrUnauthorized
	}

	tokenPair, err := model.GenerateToken(uid)
	if err != nil {
		log.Printf("fail to refresh caused by %v", err)

		return nil, ErrUnauthorized
	}

	res := &authv1.RefreshResponse{
		IdToken:      tokenPair.IDToken.String(),
		RefreshToken: tokenPair.RefreshToken.String(),
	}

	log.Printf("success to refresh")

	return connect.NewResponse(res), nil
}
