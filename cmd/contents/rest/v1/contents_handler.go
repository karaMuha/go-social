package rest

import (
	"encoding/json"
	"net/http"

	"github.com/karaMuha/go-social/contents/application/commands"
	ports "github.com/karaMuha/go-social/contents/application/ports/driver"
	"github.com/karaMuha/go-social/internal/middleware"
)

type PostsHandlerV1 struct {
	app ports.IApplication
}

func NewContentsHandlerV1(app ports.IApplication) PostsHandlerV1 {
	return PostsHandlerV1{
		app: app,
	}
}

func (h PostsHandlerV1) HandlePostContent(w http.ResponseWriter, r *http.Request) {
	var cmdParams commands.PostContentDto
	err := json.NewDecoder(r.Body).Decode(&cmdParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cmdParams.UserID = r.Context().Value(middleware.ContextUserIDKey).(string)

	_, err = h.app.PostContent(r.Context(), &cmdParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h PostsHandlerV1) HandleViewContentDetails(w http.ResponseWriter, r *http.Request) {
	postID := r.URL.Query().Get("post-id")

	post, err := h.app.GetContentDetails(r.Context(), postID)
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

func (h PostsHandlerV1) HandleUpdateContent(w http.ResponseWriter, r *http.Request) {
	var cmdParams commands.UpdateContentDto
	err := json.NewDecoder(r.Body).Decode(&cmdParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cmdParams.ID = r.PathValue("id")
	cmdParams.UserID = r.Context().Value(middleware.ContextUserIDKey).(string)

	err = h.app.UpdateContent(r.Context(), &cmdParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h PostsHandlerV1) HandleRemoveContent(w http.ResponseWriter, r *http.Request) {
	var cmdParams commands.RemoveContentDto
	err := json.NewDecoder(r.Body).Decode(&cmdParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cmdParams.UserID = r.Context().Value(middleware.ContextUserIDKey).(string)

	err = h.app.RemoveContent(r.Context(), &cmdParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
