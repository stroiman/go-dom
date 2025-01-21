// Package logger provides the basic functionality of supplying a custom logger.
package logger

import (
	"log/slog"

	"github.com/stroiman/go-dom/browser/internal/log"
)

// SetDefault sets the [slog.Logger] that will receive log messages from the
// server.
func SetDefault(logger *slog.Logger) {
	log.SetDefault(logger)
}
