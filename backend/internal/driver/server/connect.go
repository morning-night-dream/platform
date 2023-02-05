package server

import (
	"net/http"

	"github.com/bufbuild/connect-go"
	"github.com/morning-night-dream/platform/internal/adapter/handler"
	"github.com/morning-night-dream/platform/pkg/connect/proto/article/v1/articlev1connect"
	"github.com/morning-night-dream/platform/pkg/connect/proto/auth/v1/authv1connect"
	"github.com/morning-night-dream/platform/pkg/connect/proto/health/v1/healthv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func NewConnectHandler(
	health *handler.Health,
	article *handler.Article,
	auth *handler.Auth,
) http.Handler {
	interceptor := connect.WithInterceptors(NewInterceptor())

	mux := NewRouter(
		NewRoute(healthv1connect.NewHealthServiceHandler(health, interceptor)),
		NewRoute(articlev1connect.NewArticleServiceHandler(article, interceptor)),
		NewRoute(authv1connect.NewAuthServiceHandler(auth, interceptor)),
	).Mux()

	return h2c.NewHandler(mux, &http2.Server{})
}
