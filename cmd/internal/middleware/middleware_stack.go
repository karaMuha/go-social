package middleware

import (
	"net/http"

	authtoken "github.com/karaMuha/go-social/internal/auth_token"
)

type Middleware func(http.Handler, authtoken.TokenGenerator) http.Handler

func CreateStack(mws ...Middleware) Middleware {
	return func(next http.Handler, tokenGenerator authtoken.TokenGenerator) http.Handler {
		for i := len(mws) - 1; i >= 0; i-- {
			mw := mws[i]
			next = mw(next, tokenGenerator)
		}

		return next
	}
}
