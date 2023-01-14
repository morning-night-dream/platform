package handler

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"net/textproto"
	"strconv"
	"strings"
	"time"

	"github.com/bufbuild/connect-go"
	"github.com/google/uuid"
	"github.com/morning-night-dream/platform/internal/cache"
	"github.com/morning-night-dream/platform/internal/firebase"
	"github.com/morning-night-dream/platform/internal/model"
	authv1 "github.com/morning-night-dream/platform/pkg/proto/auth/v1"
)

type Auth struct {
	firebase *firebase.Client
	cache    *cache.Client
}

const age = 5 * 60

func NewAuth(firebase *firebase.Client, cache *cache.Client) *Auth {
	return &Auth{
		firebase: firebase,
		cache:    cache,
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

	// firebase に新規登録
	if err := a.firebase.CreateUser(ctx, uuid.NewString(), email, password); err != nil {
		return nil, err
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

	// firebase にログイン
	sres, err := a.firebase.Login(ctx, email, password)
	if err != nil {
		log.Printf("fail to sign in caused by %s", err)
		return nil, ErrUnauthorized
	}

	// アクセストークンからセッショントークンを取得 -> 現状はおれおれセッショントークンで対応するので不要

	// アクセストークン/リフレッシュトークン/セッショントークンを紐づけてキャッシュに保存
	exp, _ := strconv.Atoi(sres.ExpiresIn)

	strs := strings.Split(sres.IDToken, ".")

	payload, _ := base64.StdEncoding.DecodeString(strs[1])

	var mapData map[string]interface{}

	if err := json.Unmarshal(payload, &mapData); err != nil {
		return nil, err
	}

	sessionToken := uuid.NewString()

	au := model.Auth{
		ID:           sessionToken,
		UserID:       mapData["user_id"].(string),
		IDToken:      sres.IDToken,
		RefreshToken: sres.RefreshToken,
		SessionToken: sessionToken,
		ExpiresIn:    exp,
	}

	a.cache.Set(ctx, sessionToken, au)

	// セッショントークンを返す
	res := connect.NewResponse(&authv1.SignInResponse{})

	cookie := http.Cookie{
		Name:       "token",
		Value:      sessionToken,
		Path:       "",
		Domain:     "",
		Expires:    time.Now().Add(60 * time.Minute),
		RawExpires: "",
		MaxAge:     age,
		Secure:     true,
		HttpOnly:   true,
		SameSite:   0,
		Raw:        "",
		Unparsed:   []string{},
	}

	res.Header().Set("Set-Cookie", cookie.String())

	return res, nil
}

func GetToken(h http.Header) (http.Cookie, error) {
	lines := h["Cookie"]
	if len(lines) == 0 {
		return http.Cookie{}, ErrUnauthorized
	}

	for _, line := range lines {
		line = textproto.TrimString(line)

		var part string

		for len(line) > 0 { // continue since we have rest
			part, line, _ = strings.Cut(line, ";")
			part = textproto.TrimString(part)
			if part == "" {
				continue
			}
			name, val, _ := strings.Cut(part, "=")
			if name != "token" {
				return http.Cookie{}, ErrUnauthorized
			}
			return http.Cookie{Name: name, Value: val}, nil
		}
	}

	return http.Cookie{}, ErrUnauthorized
}
