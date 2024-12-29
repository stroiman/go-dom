package test

import (
	"log/slog"
	"os"
)

func FilterLogAttributes(groups []string, a slog.Attr) slog.Attr {
	// Time/level is noise in output
	if a.Key == slog.TimeKey || a.Key == slog.LevelKey {
		return slog.Attr{}
	}
	return a
}

func InstallDefaultTextLogger() *slog.LevelVar {
	var logLevel = new(slog.LevelVar)
	var h slog.Handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:       logLevel,
		ReplaceAttr: FilterLogAttributes,
	})
	slog.SetDefault(slog.New(h))
	return logLevel
}
