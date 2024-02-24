package router

import (
	"url-shortener/internal/http-server/handlers/url/save"
	mymiddleware "url-shortener/internal/http-server/my_middleware"
	"url-shortener/internal/models"

	"golang.org/x/exp/slog"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRouter(logger *slog.Logger, model *models.URL) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	router.Use(mymiddleware.New(logger))

	router.Post("/url", save.New(logger, model))

	return router
}
