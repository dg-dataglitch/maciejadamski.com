package main

import (
	"log/slog"
	"net/http"
	"os"

	"rkb/pkg/config"
	"rkb/pkg/logger"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	l, err := logger.New(cfg.LogLevel)
	if err != nil {
		l, _ = logger.New("info")
	}
	slog.SetDefault(l)

	// Serve the dist directory
	addr := cfg.Addr()
	l.Info("dev_server_listening", "addr", "http://"+addr)

	fs := http.FileServer(http.Dir(cfg.DistPath))
	http.Handle("/", fs)

	if err := http.ListenAndServe(addr, nil); err != nil {
		l.Error("server_failed", "err", err)
		os.Exit(1)
	}
}
