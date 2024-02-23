package main

import (
	"fmt"
	"log/slog"

	"url-shortener/internal/config"
	"url-shortener/logger"
)

func main() {
	cfg := config.LoadConfig()
	fmt.Println(cfg)

	logger := logger.SetupLogger(cfg.Env)
	logger = logger.With(slog.String("env", cfg.Env))
	logger.Info("starting url-shortener")
	logger.Debug("Debugging application")

	
}
