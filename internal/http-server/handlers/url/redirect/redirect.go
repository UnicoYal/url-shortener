package redirect

import (
	"fmt"
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
		op := "handlers.url.Redirect.New"

		logger = logger.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		alias := chi.URLParam(r, "alias")

		if alias == "" {
			errorMsg := "Invalid request: alias is empty"
			logger.Error(errorMsg)

			render.JSON(w, r, response.Error(errorMsg))

			return
		}

		urlToGet, err := models.GetUrl(alias)

		if err != nil {
			errorMsg := fmt.Sprintf("Cannot find url with alias: %s", alias)

			logger.Error(errorMsg)

			render.JSON(w, r, response.Error(errorMsg))

			return
		}

		logger.Info("url found, redirecting")

		http.Redirect(w, r, urlToGet.Url, http.StatusFound)
	}
}
