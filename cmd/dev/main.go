package main

import (
	"log/slog"
	"net/http"
	"os"
	"os/exec"

	"rkb/pkg/config"
	"rkb/pkg/logger"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}
	l, err := logger.New(cfg.LogLevel)
	if err != nil {
		l, _ = logger.New("info")
	}
	slog.SetDefault(l)
	l.Info("dev_server_starting", "running", "build")
	buildCmd := exec.Command("go", "run", "./cmd/build")
	buildCmd.Stdout = os.Stdout
	buildCmd.Stderr = os.Stderr
	if err := buildCmd.Run(); err != nil {
		l.Error("build_failed", "err", err)
		os.Exit(1)
	}
	addr := cfg.Addr()
	l.Info("dev_server_listening", "addr", "http://"+addr)
	fs := http.FileServer(http.Dir(cfg.DistPath))
	http.Handle("/", fs)
	if err := http.ListenAndServe(addr, nil); err != nil {
		l.Error("server_failed", "err", err)
		os.Exit(1)
	}
}
