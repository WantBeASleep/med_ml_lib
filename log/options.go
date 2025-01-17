package log

import (
	"log/slog"
	"os"
	"time"

	"github.com/fatih/color"

	"github.com/rs/zerolog"
)

type LogOption interface {
	applyOpt(cfg *config)
}

type logOption func(*config)

func (o logOption) applyOpt(c *config) {
	o(c)
}

func WithProdPreset() LogOption {
	return logOption(func(c *config) {
		c.logger = zerolog.New(os.Stdout)
		c.level = slog.LevelError
	})
}

func WithStagePreset() logOption {
	return logOption(func(c *config) {
		WithProdPreset().applyOpt(c)
		c.level = slog.LevelDebug
	})
}

func WithLocalPreset() logOption {
	return logOption(func(c *config) {
		c.logger = zerolog.New(zerolog.ConsoleWriter{
			Out: os.Stdout,
			FormatTimestamp: func(i interface{}) string {
				t, _ := time.Parse(time.DateTime, i.(string))
				return color.HiMagentaString("%s", t.Format(time.DateTime))
			},
		})
		c.level = slog.LevelDebug
	})
}
