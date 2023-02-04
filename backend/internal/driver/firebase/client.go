package firebase

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/morning-night-dream/platform/pkg/log"
	"google.golang.org/api/option"
)

// @see https://firebase.google.com/docs/reference/rest/auth
type APIRestClient struct {
	*http.Client
	Endpoint string
	APIKey   string
}

type Client struct {
	admin *auth.Client
	api   *APIRestClient
}

func NewClient(secret string, endpoint string, key string) *Client {
	opt := option.WithCredentialsJSON([]byte(secret))

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Log().Fatal("error initializing app", log.ErrorField(err))
	}

	admin, err := app.Auth(context.Background())
	if err != nil {
		log.Log().Fatal("error create auth client", log.ErrorField(err))
	}

	if endpoint == "" {
		log.Log().Fatal("error firebase api endpoint is empty")
	}

	if key == "" {
		log.Log().Fatal("error firebase api key is empty")
	}

	api := &APIRestClient{
		Client:   &http.Client{},
		Endpoint: endpoint,
		APIKey:   key,
	}

	return &Client{
		admin: admin,
		api:   api,
	}
}

func (c *Client) CreateUser(ctx context.Context, userID, email, password string) error {
	params := (&auth.UserToCreate{}).
		UID(userID).
		Email(email).
		EmailVerified(false).
		Password(password).
		Disabled(false)

	if _, err := c.admin.CreateUser(ctx, params); err != nil {
		log.GetLogCtx(ctx).Warn("faile to create user", log.ErrorField(err))

		return err
	}

	return nil
}

type SignInRequest struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	ReturnSecureToken bool   `json:"returnSecureToken"`
}

type SignInResponse struct {
	ExpiresIn    string `json:"expiresIn"`
	LocalID      string `json:"localId"`
	IDToken      string `json:"idToken"`
	RefreshToken string `json:"refreshToken"`
}

func (f *Client) Login(ctx context.Context, email, password string) (SignInResponse, error) {
	// https://firebase.google.com/docs/reference/rest/auth#section-sign-in-email-password
	url := fmt.Sprintf("%s/v1/accounts:signInWithPassword?key=%s", f.api.Endpoint, f.api.APIKey)

	req := SignInRequest{
		Email:             email,
		Password:          password,
		ReturnSecureToken: true,
	}

	var buf bytes.Buffer

	err := json.NewEncoder(&buf).Encode(req)
	if err != nil {
		log.GetLogCtx(ctx).Warn("failed to encode json", log.ErrorField(err))

		return SignInResponse{}, err
	}

	res, err := f.api.Post(url, "application/json", &buf)
	if err != nil {
		log.GetLogCtx(ctx).Warn("failed to post "+url, log.ErrorField(err))

		return SignInResponse{}, err
	}

	if res.StatusCode != http.StatusOK {
		message, err := io.ReadAll(res.Body)
		if err != nil {
			log.GetLogCtx(ctx).Warn("faile to read body", log.ErrorField(err))

			message = []byte(fmt.Sprintf("could not laod message caused by %v", err))
		}

		return SignInResponse{}, fmt.Errorf("firebase error. status code is %d, message is %v", res.StatusCode, string(message))
	}

	var response SignInResponse

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		log.GetLogCtx(ctx).Warn("failed to decode json", log.ErrorField(err))

		return SignInResponse{}, err
	}

	// 試しにfirebaseのセッショントークンを作ってみたコード
	// cookie, err := f.admin.SessionCookie(ctx, response.IDToken, time.Hour*24*5)
	// if err != nil {
	// 	return SignInResponse{}, err
	// }

	// log.Printf("session cookie: %s", cookie)

	return response, nil
}

type RefreshRequest struct {
	GrantType    string `json:"grant_type"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshResponse struct {
	ExpiresIn    string `json:"expires_in"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	IDToken      string `json:"id_token"`
	UserID       string `json:"user_id"`
	ProjectID    string `json:"project_id"`
}

func (f *Client) RefreshToken(ctx context.Context, refresh string) (string, error) {
	// https://firebase.google.com/docs/reference/rest/auth#section-refresh-token
	url := fmt.Sprintf("%s/v1/token?key=%s", f.api.Endpoint, f.api.APIKey)

	req := RefreshRequest{
		GrantType:    "refresh_token",
		RefreshToken: refresh,
	}

	var body bytes.Buffer

	if err := json.NewEncoder(&body).Encode(req); err != nil {
		return "", err
	}

	res, err := f.api.Post(url, "application/json", &body)
	if err != nil {
		return "", err
	}

	if res.StatusCode != http.StatusOK {
		message, err := io.ReadAll(res.Body)
		if err != nil {
			message = []byte(fmt.Sprintf("could not load message caused by %v", err))
		}

		return "", fmt.Errorf("firebase error. status code is %d, message is %v", res.StatusCode, string(message))
	}

	var response RefreshResponse

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return "", err
	}

	return response.IDToken, nil

}

func (f *Client) ChangePassword(ctx context.Context, uid, password string) error {
	params := (&auth.UserToUpdate{}).
		Password(password)

	if _, err := f.admin.UpdateUser(ctx, uid, params); err != nil {
		return err
	}

	return nil
}

func (f *Client) VerifyIDToken(ctx context.Context, idToken string) error {
	_, err := f.admin.VerifyIDToken(ctx, idToken)
	if err != nil {
		return err
	}

	return nil
}

func (f *Client) DeleteUser(ctx context.Context, uid string) error {
	if err := f.admin.DeleteUser(ctx, uid); err != nil {
		return err
	}

	return nil
}
