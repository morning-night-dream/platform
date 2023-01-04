package server

import (
	"net/http"

	"github.com/morning-night-dream/platform/internal/core/model"
	"github.com/morning-night-dream/platform/pkg/log"
	"github.com/newrelic/go-agent/v3/integrations/nrzap"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type Router struct {
	newrelic *newrelic.Application
	routes   []Route
}

type Route struct {
	path    string
	handler http.Handler
}

func NewRoute(path string, handler http.Handler) Route {
	return Route{
		path:    path,
		handler: handler,
	}
}

func NewRouter(routes ...Route) *Router {
	app, _ := newrelic.NewApplication(
		newrelic.ConfigAppName(model.Config.NewRelicAppName),
		newrelic.ConfigLicense(model.Config.NewRelicLicense),
		newrelic.ConfigAppLogForwardingEnabled(true),
		func(c *newrelic.Config) {
			c.Logger = nrzap.Transform(log.Log())
		},
	)

	return &Router{
		newrelic: app,
		routes:   routes,
	}
}

func (r Router) Mux() *http.ServeMux {
	mux := http.NewServeMux()

	for _, route := range r.routes {
		path := route.path
		handler := route.handler

		if model.Env.IsProd() {
			path, handler = newrelic.WrapHandle(r.newrelic, route.path, route.handler)
		}

		mux.Handle(path, handler)
	}

	return mux
}
