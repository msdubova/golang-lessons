package main

import (
	"net/http"
	"os"
	"web21/internal/api"
	"web21/internal/core"
	"web21/internal/storage"

	"github.com/rs/zerolog/log"
)

func main() {
	storage, err := storage.New(os.Getenv("POSTGRES_CONN_STR"))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to storage")
	}

	urlService := &core.ShortURLService{
		Storage: storage,
	}

	redirectHandler := &api.Redirect{
		Service: urlService,
	}

	shortURLHandler := &api.ShortURL{
		Service: urlService,
	}

	mux := http.NewServeMux()

	redirectHandler.RegisterRoutes(mux)

	shortURLHandler.RegisterRoutes(mux)

	addr := os.Getenv("ADDR")

	log.Info().Msgf("Accepting incoming requests on %v", addr)

	err = http.ListenAndServe(addr, mux)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to listen and serve")
	}
}
