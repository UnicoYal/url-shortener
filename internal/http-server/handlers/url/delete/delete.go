package delete

import (
	"net/http"
	"url-shortener/internal/models"
	"url-shortener/lib/api/response"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"golang.org/x/exp/slog"
)

func New(logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		op := "handlers.url.Delete.New"

		logger.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		alias := chi.URLParam(r, "alias")

		if alias == "" {
			logger.Error("Invalid request: Alias empty")

			render.JSON(w, r, response.Error("Empty alias"))

			return
		}

		_, err := models.DeleteUrl(alias)

		if err != nil {
			logger.Error(err.Error())

			render.JSON(w, r, response.Error(err.Error()))

			return
		}

		logger.Info("Successful deletion of url")

		render.JSON(w, r, response.OK())
	}
}
