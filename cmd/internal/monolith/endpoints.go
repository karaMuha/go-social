package monolith

var protectedEndpoints map[string]bool

func setProtectedEndpoints() {
	protectedEndpoints = make(map[string]bool)
	protectedEndpoints["POST /v1/users/signup-for-registration"] = false
	protectedEndpoints["POST  /v1/users/confirm-registration"] = false
	protectedEndpoints["GET  /v1/users/view-user-details"] = false
	protectedEndpoints["POST  /v1/users/login"] = false

	protectedEndpoints["POST /v1/followers/follow-user"] = true
	protectedEndpoints["POST /v1/followers/unfollow-user"] = true
	protectedEndpoints["GET /v1/followers/list-followers-of-user"] = true

	protectedEndpoints["POST /v1/contents/post-content"] = true
	protectedEndpoints["GET /v1/contents/view-content-details"] = true
	protectedEndpoints["POST /v1/contents/update-content"] = true
	protectedEndpoints["POST /v1/contents/remove-content"] = true
	protectedEndpoints["GET /v1/contents/view-users-content"] = false
}

func IsProtectedEndpoint(endpoint string) bool {
	return protectedEndpoints[endpoint]
}
