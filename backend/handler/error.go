package handler

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var ErrInternal = status.New(
	codes.Internal,
	"internal server",
).Err()

var ErrInvalidArgument = status.New(
	codes.InvalidArgument,
	"bad request",
).Err()

var ErrUnauthorized = status.New(
	codes.Unauthenticated,
	"unauthorized",
).Err()
