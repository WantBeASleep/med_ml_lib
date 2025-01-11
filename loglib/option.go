package loglib

import (
	"os"
	"log/slog"
)

type LopOption interface {
	applyOpt(cfg *config)
}

type logOption func(*config)

func (o logOption) applyOpt(c *config) {
	o(c)
}

func WithFileOutput(path string) LopOption {
	f, _ := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0o666)
	return logOption(func(c *config) {
		c.dest = f
	})
}

func WithDevEnv() LopOption {
	return logOption(func(c *config) {
		c.dest = os.Stdout
		c.level = slog.LevelDebug
	})
}

func WithDebugLog() LopOption {
	return logOption(func(c *config) {
		c.level = slog.LevelDebug
	})
}