package delete

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	resp "url-shortener/internal/lib/api/response"
	"url-shortener/internal/storage"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type Response struct {
	resp.Response
	Alias string `json:"alias,omitempty"`
}

type URLDeleter interface {
	DeleteURL(alias string) error
}

func New(log *slog.Logger, urlDeleter URLDeleter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.delete.New"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		alias := chi.URLParam(r, "alias")

		if alias == "" {
			log.Info("alias is empty")

			render.JSON(w, r, resp.Error("not found alias"))

			return
		}

		err := urlDeleter.DeleteURL(alias)

		fmt.Printf("DeleteURL returned: %v, type: %T\n", err, err)

		if errors.Is(err, storage.ErrURLNotFound) {
			log.Info("url not found", "alias", alias)
			render.JSON(w, r, resp.Error("not found"))
			return
		}

		if err != nil {
			log.Error("failed to delete url", slog.String("error", err.Error()))
			render.JSON(w, r, resp.Error("internal error"))
			return
		}

		render.JSON(w, r, Response{
			Response: resp.OK(),
			Alias:    alias,
		})

	}
}
