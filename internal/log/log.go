package log

import (
	"context"
	"log/slog"
)

var defaultLogger *slog.Logger

func SetDefault(logger *slog.Logger) {
	defaultLogger = logger
}

type nullHandler struct{}

func (_ nullHandler) Enabled(context.Context, slog.Level) bool  { return false }
func (_ nullHandler) Handle(context.Context, slog.Record) error { return nil }
func (_ nullHandler) WithAttrs([]slog.Attr) slog.Handler        { return nullHandler{} }
func (_ nullHandler) WithGroup(name string) slog.Handler        { return nullHandler{} }

func init() {
	defaultLogger = slog.New(nullHandler{})
}

func Info(msg string, args ...any) {
	defaultLogger.Info(msg, args...)
}

func Warn(msg string, args ...any) {
	defaultLogger.Warn(msg, args...)
}

func Debug(msg string, args ...any) {
	defaultLogger.Debug(msg, args...)
}

func Error(msg string, args ...any) {
	defaultLogger.Error(msg, args...)
}
