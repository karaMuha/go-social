package rest

import (
	"net/http"

	ports "github.com/karaMuha/go-social/posts/application/ports/driver"
)

type PostsHandlerV1 struct {
	app ports.IApplication
}

func NewPostsHandlerV1(app ports.IApplication) PostsHandlerV1 {
	return PostsHandlerV1{
		app: app,
	}
}

func (h PostsHandlerV1) PostCreationHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "endpoint not implemented yet", http.StatusNotImplemented)
}
