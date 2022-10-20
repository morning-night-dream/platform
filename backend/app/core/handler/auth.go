package handler

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"log"
	"strings"

	"github.com/bufbuild/connect-go"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/morning-night-dream/platform/app/core/database/store"
	"github.com/morning-night-dream/platform/app/core/model"
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

	return connect.NewResponse(&authv1.SignUpResponse{}), nil
}

func (a Auth) SignIn(
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

	return connect.NewResponse(res), nil
}

func (a Auth) Refresh(
	ctx context.Context,
	req *connect.Request[authv1.RefreshRequest],
) (*connect.Response[authv1.RefreshResponse], error) {
	idTokenString := req.Header().Get("Authorization")

	payload := strings.Split(idTokenString, ".")[1]

	decode, err := base64.RawURLEncoding.DecodeString(payload)
	if err != nil {
		log.Printf("fail to refresh caused by %v", err)

		return nil, ErrUnauthorized
	}

	var p struct {
		UID string `json:"uid"`
	}

	if err := json.Unmarshal(decode, &p); err != nil {
		log.Printf("fail to refresh caused by %v", err)

		return nil, ErrUnauthorized
	}

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

	if sub, _ := claims["sub"].(string); sub != p.UID {
		log.Printf("fail to refresh caused by invalid uid")

		return nil, ErrUnauthorized
	}

	tokenPair, err := model.GenerateToken(p.UID)
	if err != nil {
		log.Printf("fail to refresh caused by %v", err)

		return nil, ErrUnauthorized
	}

	res := &authv1.RefreshResponse{
		IdToken:      tokenPair.IDToken.String(),
		RefreshToken: tokenPair.RefreshToken.String(),
	}

	return connect.NewResponse(res), nil
}
