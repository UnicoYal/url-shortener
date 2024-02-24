package save

import (
	"golang.org/x/exp/slog"
	"net/http"
	"url-shortener/internal/models"
	"url-shortener/lib/api/response"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type Request struct {
	Url string `json:"url"`
	Alias string `json:"alias"`
}

type Response struct {
	response.Response
	Alias string `json:"alias,omitempty"`
}

type URLSaver interface {
	SaveURL(url string, alias string) *models.URL
}

func New(logger *slog.Logger, urlSaver URLSaver) http.HandlerFunc {
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

		newUrl := urlSaver.SaveURL(req.Url, req.Alias)

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
