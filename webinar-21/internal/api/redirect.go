package api

import (
	"context"
	"errors"
	"net/http"
	"web21/internal/core"

	"github.com/rs/zerolog/log"
)

type RedirectService interface {
	OpenShortURL(ctx context.Context, id string) (*core.ShortURL, error)
}

type Redirect struct {
	Service RedirectService
}

func (h *Redirect) RegisterRoutes(m *http.ServeMux) {
	m.HandleFunc("GET /r/{id}", h.HandleRedirect)
}

func (h *Redirect) HandleRedirect(w http.ResponseWriter, r *http.Request) {
	shortURl, err := h.Service.OpenShortURL(r.Context(), r.PathValue("id"))
	if err != nil {
		if errors.Is(err, core.ErrURLNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		log.Error().Err(err).Msg("Failed to open short URL")
		return
	}

	http.Redirect(w, r, shortURl.OriginalURL, http.StatusMovedPermanently)
}
