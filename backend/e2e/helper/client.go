package helper

import (
	"net/http"
	"testing"

	"github.com/bufbuild/connect-go"
	"github.com/morning-night-dream/platform/pkg/proto/article/v1/articlev1connect"
	"github.com/morning-night-dream/platform/pkg/proto/auth/v1/authv1connect"
	"github.com/morning-night-dream/platform/pkg/proto/health/v1/healthv1connect"
)

type Client struct {
	Article articlev1connect.ArticleServiceClient
	Health  healthv1connect.HealthServiceClient
	Auth    authv1connect.AuthServiceClient
}

func NewClient(t *testing.T, client connect.HTTPClient, url string) *Client {
	t.Helper()

	ac := articlev1connect.NewArticleServiceClient(
		client,
		url,
	)

	hc := healthv1connect.NewHealthServiceClient(
		client,
		url,
	)

	auc := authv1connect.NewAuthServiceClient(
		client,
		url,
	)

	return &Client{
		Article: ac,
		Health:  hc,
		Auth:    auc,
	}
}

func NewPlainClient(t *testing.T, url string) *Client {
	t.Helper()

	return NewClient(t, http.DefaultClient, url)
}

func NewClientWithAPIKey(t *testing.T, key string, url string) *Client {
	t.Helper()

	client := &http.Client{
		Transport: NewAPIKeyTransport(t, key),
	}

	return NewClient(t, client, url)
}

func NewClientWithCookie(t *testing.T, cookie string, url string) *Client {
	t.Helper()

	client := &http.Client{
		Transport: NewCookieTransport(t, cookie),
	}

	return NewClient(t, client, url)
}

type APIKeyTransport struct {
	t         *testing.T
	APIKey    string
	Transport http.RoundTripper
}

func NewAPIKeyTransport(
	t *testing.T,
	key string,
) *APIKeyTransport {
	return &APIKeyTransport{
		t:         t,
		APIKey:    key,
		Transport: http.DefaultTransport,
	}
}

func (at *APIKeyTransport) transport() http.RoundTripper {
	at.t.Helper()

	return at.Transport
}

func (at *APIKeyTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	at.t.Helper()

	req.Header.Add("X-API-KEY", at.APIKey)

	resp, err := at.transport().RoundTrip(req)
	if err != nil {
		return nil, err
	}

	return resp, err
}

type CookieTransport struct {
	t         *testing.T
	Cookie    string
	Transport http.RoundTripper
}

func NewCookieTransport(
	t *testing.T,
	cookie string,
) *CookieTransport {
	return &CookieTransport{
		t:         t,
		Cookie:    cookie,
		Transport: http.DefaultTransport,
	}
}

func (ct *CookieTransport) transport() http.RoundTripper {
	ct.t.Helper()

	return ct.Transport
}

func (ct *CookieTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	ct.t.Helper()

	req.Header.Add("Cookie", ct.Cookie)

	resp, err := ct.transport().RoundTrip(req)
	if err != nil {
		return nil, err
	}

	return resp, err
}
