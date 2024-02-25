package router

import (
	"url-shortener/internal/http-server/handlers/url/redirect"
	"url-shortener/internal/http-server/handlers/url/save"
	mymiddleware "url-shortener/internal/http-server/my_middleware"

	"golang.org/x/exp/slog"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRouter(logger *slog.Logger) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	router.Use(mymiddleware.New(logger))

	router.Post("/url", save.New(logger))
	router.Get("/{alias}", redirect.New(logger))

	return router
}
