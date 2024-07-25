package api

import (
	"context"
	"errors"
	"net/http"
	"web21/internal/core"
)

type UserService interface {
	GetUserByNickname(ctx context.Context, nickname string) (*core.User, error)
}

type Auth struct {
	Service UserService
}

func (a *Auth) CheckAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		user, err := a.Service.GetUserByNickname(r.Context(), username)
		if err != nil {
			if errors.Is(err, core.ErrUserNotFound) {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if user.Password != password {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "user", user)

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
