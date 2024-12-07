package monolith

import "github.com/vodkaslime/wildcard"

var ProtectedRoutes map[string]bool
var matcher *wildcard.Matcher

func SetProtectedRoutes() {
	matcher = wildcard.NewMatcher()
	ProtectedRoutes = make(map[string]bool)
	ProtectedRoutes["POST /v1/users/"] = false
}
