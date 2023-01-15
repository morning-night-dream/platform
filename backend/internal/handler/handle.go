package handler

import (
	"context"
	"net/http"
	"net/textproto"
	"strings"

	"github.com/morning-night-dream/platform/internal/cache"
	"github.com/morning-night-dream/platform/internal/firebase"
	"github.com/morning-night-dream/platform/internal/model"
)

type Handle struct {
	firebase *firebase.Client
	cache    *cache.Client
}

func NewHandle(
	firebase *firebase.Client,
	cache *cache.Client,
) *Handle {
	return &Handle{
		firebase: firebase,
		cache:    cache,
	}
}

func (h *Handle) Authorize(ctx context.Context, header http.Header) (model.Auth, error) {
	cookie, err := h.getToken(header)
	if err != nil {
		return model.Auth{}, ErrUnauthorized
	}

	auth, err := h.cache.Get(ctx, cookie.Value)
	if err != nil {
		return model.Auth{}, ErrUnauthorized
	}

	if err := h.firebase.VerifyIDToken(ctx, auth.IDToken); err != nil {
		return model.Auth{}, ErrUnauthorized
	}

	return auth, nil
}

func (h *Handle) getToken(header http.Header) (http.Cookie, error) {
	lines := header["Cookie"]
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

func (h *Handle) GetSession(header http.Header) (string, error) {
	lines := header["Cookie"]
	if len(lines) == 0 {
		return "", ErrUnauthorized
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
				return "", ErrUnauthorized
			}
			return val, nil
		}
	}

	return "", ErrUnauthorized
}
