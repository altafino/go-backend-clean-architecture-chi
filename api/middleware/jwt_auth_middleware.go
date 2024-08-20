package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/altafino/go-backend-clean-architecture-chi/internal/tokenutil"
)

func JwtAuthMiddleware(secret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			t := strings.Split(authHeader, " ")
			if len(t) == 2 {
				authToken := t[1]
				authorized, err := tokenutil.IsAuthorized(authToken, secret)
				if authorized {
					userID, err := tokenutil.ExtractIDFromToken(authToken, secret)
					if err != nil {
						http.Error(w, jsonError(err.Error()), http.StatusUnauthorized)
						return
					}
					ctx := context.WithValue(r.Context(), "x-user-id", userID)
					next.ServeHTTP(w, r.WithContext(ctx))
					return
				}
				http.Error(w, jsonError(err.Error()), http.StatusUnauthorized)
				return
			}
			http.Error(w, jsonError("Not authorized"), http.StatusUnauthorized)
		})
	}
}

func jsonError(message string) string {
	return `{"message": "` + message + `"}`
}
