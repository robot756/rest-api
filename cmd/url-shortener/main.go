package main

import (
	"log/slog"
	"os"
	"rest-api/internal/config"
)

func main() {
	// TODO: init config: cleanevn
	cfg := config.MustLoad()

	// TODO: init logger: slog
	logger := setupLogger(cfg.ENV)

	logger.Info("starting url-shortener")
	logger.Debug("debug message are enable")

	// TODO: init storage: postgreSQL

	// TODO: init router: chi

	// TODO: init run server:
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case "local":
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	case "dev":
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	case "prod":
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
