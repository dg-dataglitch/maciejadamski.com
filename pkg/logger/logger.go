package logger

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
)

func New(logLevel string) (*slog.Logger, error) {
	var slogLevel slog.Level
	switch strings.ToLower(logLevel) {
	case "debug":
		slogLevel = slog.LevelDebug
	case "info":
		slogLevel = slog.LevelInfo
	case "warn":
		slogLevel = slog.LevelWarn
	case "error":
		slogLevel = slog.LevelError
	default:
		return nil, fmt.Errorf("invalid value for Log level: %s", logLevel)
	}
	opts := &slog.HandlerOptions{Level: slogLevel}
	handler := slog.NewTextHandler(os.Stdout, opts)
	return slog.New(handler), nil
}
