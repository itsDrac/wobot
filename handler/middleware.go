package handler

import (
	"context"
	"net/http"
	"strings"
)

func (h ChiHandler) AuthUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the request is authenticated
		if r.Header.Get("Authorization") == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		// Extract the token from the Authorization header
		split := strings.Split(r.Header.Get("Authorization"), " ")
		if len(split) != 2 || split[0] != "Bearer" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		token := split[1]
		user, err := h.Service.User.Authenticate(r.Context(), token)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		// Set the user in the request context
		ctx := context.WithValue(r.Context(), "user", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
