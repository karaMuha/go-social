package rest

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/karaMuha/go-social/posts/application/commands"
	ports "github.com/karaMuha/go-social/posts/application/ports/driver"
	"github.com/karaMuha/go-social/posts/application/queries"
)

type PostsHandlerV1 struct {
	app ports.IApplication
}

func NewPostsHandlerV1(app ports.IApplication) PostsHandlerV1 {
	return PostsHandlerV1{
		app: app,
	}
}

func (h PostsHandlerV1) HandleCreatePost(w http.ResponseWriter, r *http.Request) {
	var cmdParams commands.CreatePostDto
	err := json.NewDecoder(r.Body).Decode(&cmdParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.app.CreatePost(r.Context(), &cmdParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h PostsHandlerV1) HandleGetPost(w http.ResponseWriter, r *http.Request) {
	postID := r.PathValue("id")
	queryParams := queries.GetPostDto{
		PostID: postID,
	}

	post, err := h.app.GetPost(r.Context(), &queryParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	responseJson, err := json.Marshal(post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJson)
}

func (h PostsHandlerV1) HandleUpdatePost(w http.ResponseWriter, r *http.Request) {
	var cmdParams commands.UpdatePostDto
	err := json.NewDecoder(r.Body).Decode(&cmdParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cmdParams.ID = r.PathValue("id")
	cmdParams.UserID = ""

	err = h.app.UpdatePost(context.Background(), &cmdParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
