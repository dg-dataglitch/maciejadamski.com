package main

import (
	"log/slog"
	"os"

	"maciejadamski/pkg/server"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	setupLogger()
	port := os.Getenv("PORT")
	if err := server.Run(port, "dist"); err != nil {
		slog.Error("server_failed", "error", err)
		os.Exit(1)
	}
}

func setupLogger() {
	level := slog.LevelInfo
	if env := os.Getenv("LOG_LEVEL"); env != "" {
		_ = level.UnmarshalText([]byte(env))
	}
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	})))
}
