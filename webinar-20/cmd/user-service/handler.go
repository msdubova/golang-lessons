package main

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

func newUserHandler(s *userService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := json.NewEncoder(w).Encode(s.storage.GetAllUsers())
		if err != nil {
			log.Warn().Err(err).Msg("Failed to JSON encode")
		}
	})
}
