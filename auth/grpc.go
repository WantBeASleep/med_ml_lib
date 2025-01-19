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

// TODO: описать в доке всю эту схемку с аутентификацией
const (
	// headers
	requestIDHeader = "x-request_id"
)

func AuthServerCall(
	ctx context.Context,
	req any,
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp any, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		slog.InfoContext(ctx, "Server call w/o x-request_id", slog.String("called method", info.FullMethod))
		return nil, status.Error(codes.Unauthenticated, "x-request_id required")
	}

	requestIDArr := md.Get(requestIDHeader)
	if len(requestIDArr) != 1 {
		slog.InfoContext(ctx, "Server call w/o x-request_id", slog.String("called method", info.FullMethod))
		return nil, status.Error(codes.Unauthenticated, "x-request_id required")
	}

	requestID, err := uuid.Parse(requestIDArr[0])
	if err != nil {
		slog.InfoContext(ctx, "Server call w/o x-request_id", slog.String("called method", info.FullMethod))
		return nil, status.Error(codes.Unauthenticated, "x-request_id invalid")
	}

	ctx = withValue(ctx, request_id, requestID.String())

	return handler(ctx, req)
}

// возьмет информацию из auth ctx.
// если нет нужных ключей спаникует
func AuthEnrichClientCall(
	ctx context.Context,
	method string,
	req, reply any,
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	md := metadata.New(map[string]string{
		requestIDHeader: ctx.Value(request_id).(string),
	})

	ctx = metadata.NewOutgoingContext(ctx, md)

	return invoker(ctx, method, req, reply, cc, opts...)
}
