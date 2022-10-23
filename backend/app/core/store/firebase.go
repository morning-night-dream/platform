package store

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/morning-night-dream/platform/app/core/model"
	"google.golang.org/api/option"
)

const perm = 0o600

// @see https://firebase.google.com/docs/reference/rest/auth
type APIRestClient struct {
	*http.Client
	Endpoint string
	APIKey   string
}

type FirebaseClient struct {
	admin *auth.Client
	api   *APIRestClient
}

type secret struct {
	Type                    string `json:"type"`
	ProjectID               string `json:"project_id"`
	PrivateKeyID            string `json:"private_key_id"`
	PrivateKey              string `json:"private_key"`
	ClientEmail             string `json:"client_email"`
	ClientID                string `json:"client_id"`
	AuthURL                 string `json:"auth_url"`
	TokenURL                string `json:"token_url"`
	AuthProviderX509CertURL string `json:"auth_provider_x509_cert"`
	ClientX509CertURL       string `json:"client_x509_cert"`
}

func NewFirebaseClient() *FirebaseClient {
	s := secret{
		Type:                    os.Getenv("FIREBASE_TYPE"),
		ProjectID:               os.Getenv("FIREBASE_PROJECT_ID"),
		PrivateKeyID:            os.Getenv("FIREBASE_PRIVATE_KEY_ID"),
		PrivateKey:              os.Getenv("FIREBASE_PRIVATE_KEY"),
		ClientEmail:             os.Getenv("FIREBASE_CLIENT_EMAIL"),
		ClientID:                os.Getenv("FIREBASE_CLIENT_ID"),
		AuthURL:                 os.Getenv("FIREBASE_AUTH_URL"),
		TokenURL:                os.Getenv("FIREBASE_TOKEN_URL"),
		AuthProviderX509CertURL: os.Getenv("FIREBASE_AUTH_PROVIDER_X509_CERT_URL"),
		ClientX509CertURL:       os.Getenv("FIREBASE_CLIENT_X509_CERT_URL"),
	}

	file, _ := json.MarshalIndent(s, "", "")

	if err := os.WriteFile("secret.json", file, perm); err != nil {
		log.Fatalf("error write file secret.json: %v", err)
	}

	opt := option.WithCredentialsFile("./secret.json")

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v", err)
	}

	admin, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error create auth client: %v", err)
	}

	if err := os.Remove("./secret.json"); err != nil {
		log.Fatalf("error delete file secret.json: %v", err)
	}

	endpoint := os.Getenv("FIREBASE_API_ENDPOINT")

	if endpoint == "" {
		log.Fatalf("error firebase api endpoint is empty")
	}

	key := os.Getenv("FIREBASE_API_KEY")

	if key == "" {
		log.Fatalf("error firebase api key is empty")
	}

	api := &APIRestClient{
		Client:   &http.Client{},
		Endpoint: endpoint,
		APIKey:   key,
	}

	return &FirebaseClient{
		admin: admin,
		api:   api,
	}
}

func (c *FirebaseClient) CreateUser(ctx context.Context, userID, email, password string) error {
	params := (&auth.UserToCreate{}).
		UID(userID).
		Email(email).
		EmailVerified(false).
		Password(password).
		Disabled(false)

	_, err := c.admin.CreateUser(ctx, params)
	if err != nil {
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

func (f *FirebaseClient) Login(ctx context.Context, email, password string) (model.Tokens, error) {
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
		return model.Tokens{}, err
	}

	res, err := f.api.Post(url, "application/json", &buf)
	if err != nil {
		return model.Tokens{}, err
	}

	if res.StatusCode != http.StatusOK {
		return model.Tokens{}, fmt.Errorf("firebase error")
	}

	var response SignInResponse

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return model.Tokens{}, err
	}

	return model.Tokens{
		IDToken:      response.IDToken,
		RefreshToken: response.RefreshToken,
	}, nil
}

type RefreshRequest struct {
	Token             string `json:"token"`
	ReturnSecureToken bool   `json:"returnSecureToken"`
}

type RefreshResponse struct {
	ExpiresIn    string `json:"expiresIn"`
	IDToken      string `json:"idToken"`
	RefreshToken string `json:"refreshToken"`
}

func (f *FirebaseClient) RefreshToken(ctx context.Context, token string) (model.Tokens, error) {
	// https://firebase.google.com/docs/reference/rest/auth#section-refresh-token
	url := fmt.Sprintf("%s/v1/accounts:signInWithCustomToken?key=%s", f.api.Endpoint, f.api.APIKey)

	req := RefreshRequest{
		Token:             token,
		ReturnSecureToken: true,
	}

	var buf bytes.Buffer

	err := json.NewEncoder(&buf).Encode(req)
	if err != nil {
		return model.Tokens{}, err
	}

	res, err := f.api.Post(url, "application/json", &buf)
	if err != nil {
		return model.Tokens{}, err
	}

	if res.StatusCode != http.StatusOK {
		return model.Tokens{}, fmt.Errorf("firebase error")
	}

	var response RefreshResponse

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return model.Tokens{}, err
	}

	return model.Tokens{
		IDToken:      response.IDToken,
		RefreshToken: response.RefreshToken,
	}, nil

}

func (f *FirebaseClient) ChangePassword(ctx context.Context, uid, password string) error {
	params := (&auth.UserToUpdate{}).
		Password(password)

	if _, err := f.admin.UpdateUser(ctx, uid, params); err != nil {
		return err
	}

	return nil
}

func (f *FirebaseClient) VerifyIDToken(ctx context.Context, accessToken string) error {
	_, err := f.admin.VerifyIDToken(ctx, accessToken)
	if err != nil {
		return err
	}

	return nil
}

func (f *FirebaseClient) DeleteUser(ctx context.Context, uid string) error {
	if err := f.admin.DeleteUser(ctx, uid); err != nil {
		return err
	}

	return nil
}
