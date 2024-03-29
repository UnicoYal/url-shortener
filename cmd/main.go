package main

import (
	"fmt"

	"url-shortener/internal/config"
	"url-shortener/internal/http-server/server"
	"url-shortener/logger"
	"url-shortener/router"
)

func main() {
	cfg := config.LoadConfig()
	fmt.Println(cfg)

	logger := logger.SetupLogger(cfg.Env)

	router := router.SetupRouter(cfg, logger)

	logger.Info("Server is starting on address: %s", cfg.HTTPServer.Address)

	server.StartServer(logger, router, cfg)
}
