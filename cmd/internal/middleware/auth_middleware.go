package middleware

import (
	"context"
	"net/http"

	authtoken "github.com/karaMuha/go-social/internal/auth_token"
	"github.com/karaMuha/go-social/internal/monolith"
)

type contextUserID string

const ContextUserIDKey contextUserID = "userID"

func AuthMiddleware(next http.Handler, tokenProvider authtoken.ITokenProvider) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestTarget := r.Method + " " + r.URL.Path

		if !monolith.IsProtectedRoute(requestTarget) {
			next.ServeHTTP(w, r)
			return
		}

		accessToken, err := r.Cookie("access_token")

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		parsedToken, err := tokenProvider.VerifyToken(accessToken.Value)

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// extract user id from token for further usage
		userID, err := tokenProvider.GetUserIDFromToken(parsedToken)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		ctx := context.WithValue(r.Context(), ContextUserIDKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
