package middlewares

import (
	"context"
	"net/http"
	"server/internal/common"
)

type ContextKey string

const userIDKey ContextKey = "userID"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// gets the token from the cookie
		cookie, err := r.Cookie("token")
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// gets the user id from the token
		userID, err := common.GetUserIdFromToken(cookie.Value)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// sets the user id in the context
		ctx := context.WithValue(r.Context(), userIDKey, userID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
