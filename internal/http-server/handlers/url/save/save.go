package save

import (
	"net/http"
	"url-shortener/internal/models"
	"url-shortener/lib/api/response"

	"golang.org/x/exp/slog"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type Request struct {
	Url   string `json:"url"`
	Alias string `json:"alias"`
}

type Response struct {
	response.Response
	Alias string `json:"alias,omitempty"`
}

func New(logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.save.New"

		logger = logger.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req Request

		err := render.DecodeJSON(r.Body, &req)

		if err != nil {
			logger.Error("failed to decode request body: %s", err)

			render.JSON(w, r, response.Error(err.Error()))

			return
		}

		logger.Info("request body decoded", slog.Any("request", req))

		newUrl := models.SaveURL(req.Url, req.Alias)

		logger.Info("url added", slog.String("url", newUrl.Url))

		responseOK(w, r, newUrl.Alias)
	}
}

func responseOK(w http.ResponseWriter, r *http.Request, alias string) {
	render.JSON(w, r, Response{
		Response: response.OK(),
		Alias:    alias,
	})
}
