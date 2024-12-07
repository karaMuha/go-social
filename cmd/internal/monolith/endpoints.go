package monolith

import "github.com/vodkaslime/wildcard"

var ProtectedRoutes map[string]bool
var matcher *wildcard.Matcher

func setProtectedRoutes() {
	matcher = wildcard.NewMatcher()
	ProtectedRoutes = make(map[string]bool)
	ProtectedRoutes["POST /v1/users/"] = false
	ProtectedRoutes["PUT  /v1/users/confirm"] = false
	ProtectedRoutes["GET  /v1/users/email/*"] = false
	ProtectedRoutes["GET  /v1/users/*"] = false
	ProtectedRoutes["POST  /v1/users/login"] = false

	ProtectedRoutes["POST /v1/followers/"] = true
	ProtectedRoutes["DELETE /v1/followers/"] = true
	ProtectedRoutes["GET /v1/followers/*"] = true

	ProtectedRoutes["POST /v1/posts/"] = true
	ProtectedRoutes["GET /v1/posts/*"] = true
	ProtectedRoutes["PUT /v1/posts/*"] = true
	ProtectedRoutes["DELETE /v1/posts/*"] = true
}

func IsProtectedRoute(endpoint string) bool {
	for i, v := range ProtectedRoutes {
		if result, _ := matcher.Match(i, endpoint); result {
			return v
		}
	}

	return false
}
