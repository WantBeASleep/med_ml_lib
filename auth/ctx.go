package auth

import (
	"context"

	"github.com/google/uuid"
)

type key struct{}
type authValues map[string]any

var authKey key

func withValue(ctx context.Context, k string, v any) context.Context {
	ctxValues, ok := ctx.Value(authKey).(authValues)
	if ctxValues == nil || !ok {
		return context.WithValue(ctx, authKey, authValues{k: v})
	}

	ctxValues[k] = v
	return context.WithValue(ctx, authKey, ctxValues)
}

func getValue(ctx context.Context, k string) any {
	ctxValues, ok := ctx.Value(authKey).(authValues)
	if ctxValues == nil || !ok {
		return nil
	}

	return ctxValues[k]
}


// WITH

func WithRequestID(ctx context.Context, id uuid.UUID) context.Context {
	return withValue(ctx, request_id, id.String())
}