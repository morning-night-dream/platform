package client

import (
	"net/http"

	"github.com/morning-night-dream/platform/internal/controller"
	"github.com/morning-night-dream/platform/pkg/proto/article/v1/articlev1connect"
	"github.com/morning-night-dream/platform/pkg/proto/auth/v1/authv1connect"
	"github.com/morning-night-dream/platform/pkg/proto/health/v1/healthv1connect"
)

var _ controller.ClientFactory = (*Client)(nil)

type Client struct{}

func New() *Client {
	return &Client{}
}

func (c *Client) Of(url string) (*controller.Client, error) {
	client := http.DefaultClient

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

	return &controller.Client{
		Article: ac,
		Health:  hc,
		Auth:    auc,
	}, nil
}
