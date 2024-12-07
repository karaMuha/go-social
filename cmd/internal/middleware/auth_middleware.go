package middleware

import (
	"context"
	"net/http"

	"github.com/golang-jwt/jwt"
	authtoken "github.com/karaMuha/go-social/internal/auth_token"
	"github.com/karaMuha/go-social/internal/monolith"
)

type contextUserID string

const ContextUserIDKey contextUserID = "userID"

func AuthMiddleware(next http.Handler, tokenGenerator authtoken.TokenGenerator) http.Handler {
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

		verifiedToken, err := tokenGenerator.VerifyToken(accessToken.Value)

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// extract user id from token for further usage
		claims, ok := verifiedToken.Claims.(jwt.MapClaims)

		if !ok {
			http.Error(w, "Could not convert jwt claims", http.StatusInternalServerError)
			return
		}

		userID, ok := claims["sub"].(string)

		if !ok {
			http.Error(w, "Could not convert user id from jwt claims to string", http.StatusInternalServerError)
			return
		}

		ctx := context.WithValue(r.Context(), ContextUserIDKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
