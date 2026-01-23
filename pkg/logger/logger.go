package logger

import (
	"fmt"
	"log/slog"
	"os"
)

func New(logLevel string) (*slog.Logger, error) {
	var level slog.Level
	if err := level.UnmarshalText([]byte(logLevel)); err != nil {
		return nil, fmt.Errorf("invalid log level %q: %w", logLevel, err)
	}

	opts := &slog.HandlerOptions{Level: level}
	handler := slog.NewTextHandler(os.Stdout, opts)
	return slog.New(handler), nil
}
