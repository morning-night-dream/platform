package server

import (
	"net/http"
	"time"

	"github.com/morning-night-dream/platform/pkg/log"
	"go.uber.org/zap"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()

		ctx := log.SetLogCtx(r.Context())

		logger := log.GetLogCtx(ctx)

		next.ServeHTTP(w, r)

		logger.Info(
			"access log",
			zap.String("method", r.Method),
			zap.String("path", r.RequestURI),
			// zap.String("addr", req.Peer().Addr),
			zap.String("ua", r.Header["User-Agent"][0]),
			// zap.String("code", status.Code(err).String()),
			zap.String("elapsed", time.Since(now).String()),
			zap.Int64("elapsed(ns)", time.Since(now).Nanoseconds()),
		)
	})
}
