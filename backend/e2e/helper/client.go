package helper

import (
	"net/http"
	"testing"

	"github.com/bufbuild/connect-go"
	"github.com/morning-night-dream/platform/pkg/proto/article/v1/articlev1connect"
	"github.com/morning-night-dream/platform/pkg/proto/health/v1/healthv1connect"
)

type Client struct {
	Article articlev1connect.ArticleServiceClient
	Health  healthv1connect.HealthServiceClient
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

	return &Client{
		Article: ac,
		Health:  hc,
	}
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
