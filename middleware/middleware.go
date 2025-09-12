package middleware

import (
	"authorizate/config"
	"authorizate/pkg/jwt"
	"context"
	"net/http"
	"strings"
)

type contextKey string

const PhoneContextKey contextKey = "phone"

func AuthMiddleware(config *config.Config) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if !strings.HasPrefix(authHeader, "Bearer ") {
				WriteUnauthorized(w)
				return
			}

			token := strings.TrimPrefix(authHeader, "Bearer ")
			isValid, jwtData := jwt.NewJwt(config.Numb.Secret).Parse(token)

			if !isValid {
				WriteUnauthorized(w)
				return
			}

			ctx := context.WithValue(r.Context(), PhoneContextKey, jwtData.Phone)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func WriteUnauthorized(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte("Unauthorized"))
}
