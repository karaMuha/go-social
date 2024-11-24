package rest

import (
	"encoding/json"
	"net/http"

	"github.com/karaMuha/go-social/posts/application/commands"
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
	var requestBody commands.CreatePostDto
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.app.CreatePost(r.Context(), &requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
