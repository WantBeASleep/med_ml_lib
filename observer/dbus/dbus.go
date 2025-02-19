package dbus

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/WantBeASleep/med_ml_lib/dbus"
	"github.com/WantBeASleep/med_ml_lib/observer/consts"
	"github.com/WantBeASleep/med_ml_lib/observer/cross"
	loglib "github.com/WantBeASleep/med_ml_lib/observer/log"
	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
)

func CrossEventProduce(
	ctx context.Context,
	topic string,
	_ proto.Message,
	cfg *dbus.OptionalSendCfg,
) error {
	crossValues := cross.GetContextAttrs(ctx)
	if cfg.Headers == nil {
		cfg.Headers = map[string]string{}
	}

	for k, v := range crossValues {
		cfg.Headers[k] = fmt.Sprintf("%v", v)
	}

	return nil
}

func CrossEventConsume(
	ctx context.Context,
	stats dbus.EventStats[proto.Message],
) (context.Context, error) {
	crossValues := cross.GetContextAttrs(ctx)
	for k, v := range stats.Headers {
		crossValues[k] = v
	}
	ctx = cross.WithFields(ctx, crossValues)

	return ctx, nil
}

func LogEventProduce(
	ctx context.Context,
	topic string,
	_ proto.Message,
	_ *dbus.OptionalSendCfg,
) error {
	slog.InfoContext(ctx, "Send event", slog.String("topic", topic))
	return nil
}

func LogEventConsume(
	ctx context.Context,
	stats dbus.EventStats[proto.Message],
) (context.Context, error) {
	requestIDField, ok := stats.Headers[consts.RequestID]
	if !ok {
		slog.WarnContext(ctx, "Consume event w/o requestID", slog.String("topic", stats.Topic))
		return ctx, nil
	}

	requestID, err := uuid.Parse(requestIDField)
	if err != nil {
		slog.WarnContext(ctx, "Consume event with wrong requestID", slog.String("topic", stats.Topic), slog.String("requestID", requestIDField))
		return ctx, nil
	}

	ctx = loglib.WithFields(ctx, map[string]any{
		consts.RequestID:     requestID,
		consts.RequestMethod: stats.Topic,
	})
	slog.InfoContext(ctx, "Consume event")
	return ctx, nil
}
