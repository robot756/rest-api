package main

import (
	"fmt"
	"log/slog"
	"os"
	"url-shortener/internal/config"
)

const (
	envLocal = "local"
	envProd  = "prod"
	envDev   = "development"
)

func main() {
	// TODO: init config: cleanenv
	cfg := config.MustLoad()
	fmt.Println(cfg)

	// TODO: init logger: slog
	log := setupLogger(cfg.Env)

	log.Info("starting url-shortener")
	log.Debug("debug message are enabled")

	// TODO: init storage: postgresSql

	// TODO: init router: chi

	// TODO: run server: Docker

}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
