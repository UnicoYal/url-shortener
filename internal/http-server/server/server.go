package server

import (
	"net/http"
	"url-shortener/internal/config"

	"github.com/go-chi/chi/v5"
	"golang.org/x/exp/slog"
)

func StartServer(logger *slog.Logger, router *chi.Mux, cfg *config.Config) {
	logger.Info("Server is starting on address: %s", cfg.HTTPServer.Address)

	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.Idle_timeoute,
	}

	if err := srv.ListenAndServe(); err != nil {
		logger.Error("failed to start")
	}
}
