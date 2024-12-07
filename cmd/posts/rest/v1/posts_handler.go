package rest

import (
	"encoding/json"
	"net/http"

	"github.com/karaMuha/go-social/internal/middleware"
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

func (h PostsHandlerV1) HandleCreatePost(w http.ResponseWriter, r *http.Request) {
	var cmdParams commands.CreatePostDto
	err := json.NewDecoder(r.Body).Decode(&cmdParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cmdParams.UserID = r.Context().Value(middleware.ContextUserIDKey).(string)

	err = h.app.CreatePost(r.Context(), &cmdParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h PostsHandlerV1) HandleGetPost(w http.ResponseWriter, r *http.Request) {
	postID := r.PathValue("id")

	post, err := h.app.GetPost(r.Context(), postID)
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
	cmdParams.UserID = r.Context().Value(middleware.ContextUserIDKey).(string)

	err = h.app.UpdatePost(r.Context(), &cmdParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h PostsHandlerV1) HandleDeletePost(w http.ResponseWriter, r *http.Request) {
	var cmdParams commands.DeletePostDto
	cmdParams.ID = r.PathValue("id")
	cmdParams.UserID = r.Context().Value(middleware.ContextUserIDKey).(string)

	err := h.app.DeletePost(r.Context(), &cmdParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
