package api

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"web21/internal/core"

	"github.com/rs/zerolog/log"
)

type ShortURLService interface {
	GetShortURLByID(ctx context.Context, id string) (*core.ShortURL, error)
	CreateShortURL(ctx context.Context, user core.User, url string) (*core.ShortURL, error)
}

type ShortURL struct {
	Service ShortURLService
}

func (h *ShortURL) RegisterRoutes(m *http.ServeMux, a *Auth) {
	m.HandleFunc("GET /short_urls/{id}", a.CheckAuth(h.GetShortURLByID))
	m.HandleFunc("POST /short_urls", a.CheckAuth(h.PostShortURL))
}

func (h *ShortURL) GetShortURLByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	shortURL, err := h.Service.GetShortURLByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, core.ErrURLNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		log.Error().Err(err).Msg("Failed to get short URL")
		return
	}

	if err := json.NewEncoder(w).Encode(shortURL); err != nil {
		log.Warn().Err(err).Msg("Failed to encode JSON")
	}
}

type PostShortURLReqBody struct {
	URL string
}

func (h *ShortURL) PostShortURL(w http.ResponseWriter, r *http.Request) {
	var reqBody PostShortURLReqBody

	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user := r.Context().Value("user").(core.User)

	shortURL, err := h.Service.CreateShortURL(r.Context(), user, reqBody.URL)
	if err != nil {
		if errors.Is(err, core.ErrInvalidURL) {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]any{"error": err})
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		log.Error().Err(err).Msg("Failed to create short URL")
		return
	}

	if err := json.NewEncoder(w).Encode(shortURL); err != nil {
		log.Warn().Err(err).Msg("Failed to encode JSON")
	}
}
