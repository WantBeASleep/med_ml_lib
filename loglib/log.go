// TODO: сделать через zerolog
package loglib

import (
	"context"
	"io"
	"log/slog"
	"os"

	"github.com/WantBeASleep/goooool/ctxlib"
)

type config struct {
	dest  io.Writer
	level slog.Level
}

func defaultConfig() config {
	return config{dest: os.Stdout}
}

type handler struct {
	slog.Handler
}

func (h handler) Handle(ctx context.Context, r slog.Record) error {
	attr := ctxlib.PublicGetAll(ctx)

	contextAttrs := make([]any, 0, len(attr))
	for k, v := range attr {
		contextAttrs = append(contextAttrs, slog.Any(k, v))
	}

	r.AddAttrs(slog.Group("context", contextAttrs...))

	return h.Handler.Handle(ctx, r)
}

func InitLogger(opts ...LopOption) {
	cfg := defaultConfig()
	for _, o := range opts {
		o.applyOpt(&cfg)
	}

	jsonHandler := slog.NewJSONHandler(cfg.dest, &slog.HandlerOptions{Level: cfg.level})

	logger := slog.New(handler{Handler: jsonHandler})
	slog.SetDefault(logger)
}
