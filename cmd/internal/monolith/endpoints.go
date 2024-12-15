package monolith

import "github.com/vodkaslime/wildcard"

var ProtectedRoutes map[string]bool
var matcher *wildcard.Matcher

func setProtectedRoutes() {
	matcher = wildcard.NewMatcher()
	ProtectedRoutes = make(map[string]bool)
	ProtectedRoutes["POST /v1/users/signup-for-registration"] = false
	ProtectedRoutes["POST  /v1/users/confirm-registration"] = false
	ProtectedRoutes["GET  /v1/users/view-user-details"] = false
	ProtectedRoutes["POST  /v1/users/login"] = false

	ProtectedRoutes["POST /v1/followers/follow-user"] = true
	ProtectedRoutes["POST /v1/followers/unfollow-user"] = true
	ProtectedRoutes["GET /v1/followers/list-followers-of-user"] = true

	ProtectedRoutes["POST /v1/contents/post-content"] = true
	ProtectedRoutes["GET /v1/contents/view-content-details"] = true
	ProtectedRoutes["POST /v1/contents/update-content"] = true
	ProtectedRoutes["POST /v1/contents/remove-content"] = true
	ProtectedRoutes["GET /v1/contents/view-users-content"] = false
}

func IsProtectedRoute(endpoint string) bool {
	for i, v := range ProtectedRoutes {
		if result, _ := matcher.Match(i, endpoint); result {
			return v
		}
	}

	return false
}
