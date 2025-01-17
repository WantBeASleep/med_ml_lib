package grpclib

import (
	"context"
	"log/slog"

	"github.com/WantBeASleep/med_ml_lib/log"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	requestIdKey = "x-request_id"
	methodKey    = "x-method"
)

// Не проверяет на отсутствие метаданных. Вызовет **панику**!
func MustServerCall(
	ctx context.Context,
	req any,
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp any, err error) {
	md, _ := metadata.FromIncomingContext(ctx)
	requestID := uuid.MustParse(md.Get(requestIdKey)[0])

	ctx = log.WithFields(ctx, map[string]any{
		requestIdKey: requestID,
		methodKey: info.FullMethod,
	})
	slog.InfoContext(ctx, "Server call")

	return handler(ctx, req)
}

func ClientCall(
	ctx context.Context,
	method string,
	req, reply any,
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	slog.InfoContext(ctx, "Client call", slog.String("method", method))
	return invoker(ctx, method, req, reply, cc, opts...)
}