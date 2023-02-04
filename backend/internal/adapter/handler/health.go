package handler

import (
	"context"

	"github.com/bufbuild/connect-go"
	healthv1 "github.com/morning-night-dream/platform/pkg/proto/health/v1"
)

type Health struct{}

func NewHealth() *Health {
	return &Health{}
}

func (h Health) Check(
	ctx context.Context,
	req *connect.Request[healthv1.CheckRequest],
) (*connect.Response[healthv1.CheckResponse], error) {
	return connect.NewResponse(&healthv1.CheckResponse{}), nil
}
