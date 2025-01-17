package auth

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type key string

// TODO: описать в доке всю эту схемку с аутентификацией
const (
	requestIDKey key    = "x-request_id"
	methodHeader    string = "x-method"
)

func AuthServerCall(
	ctx context.Context,
	req any,
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp any, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		slog.InfoContext(ctx, "Info call w/o x-request_id", slog.String(methodHeader, info.FullMethod))
		return nil, status.Error(codes.Unauthenticated, "x-request_id required")
	}

	requestIDArr := md.Get(string(requestIDKey))
	if len(requestIDArr) != 1 {
		slog.InfoContext(ctx, "Info call w/o x-request_id", slog.String(methodHeader, info.FullMethod))
		return nil, status.Error(codes.Unauthenticated, "x-request_id required")
	}

	requestID, err := uuid.Parse(requestIDArr[0])
	if err != nil {
		slog.InfoContext(ctx, "Info call w/o x-request_id", slog.String(methodHeader, info.FullMethod))
		return nil, status.Error(codes.Unauthenticated, "x-request_id invalid")
	}

	ctx = context.WithValue(ctx, requestIDKey, requestID.String())

	return handler(ctx, req)
}

func AuthEnrichClientCall(
	ctx context.Context,
	method string,
	req, reply any,
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	md := metadata.New(map[string]string{
		string(requestIDKey): ctx.Value(requestIDKey).(string),
	})

	ctx = metadata.NewOutgoingContext(ctx, md)

	return invoker(ctx, method, req, reply, cc, opts...)
}
