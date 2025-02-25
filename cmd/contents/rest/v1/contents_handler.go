package rest

import (
	"encoding/json"
	"net/http"

	"github.com/karaMuha/go-social/contents/application/commands"
	ports "github.com/karaMuha/go-social/contents/application/ports/driver"
	"github.com/karaMuha/go-social/internal/http/response"
	"github.com/karaMuha/go-social/internal/middleware"
)

type ContentsHandlerV1 struct {
	app ports.IApplication
}

func NewContentsHandlerV1(app ports.IApplication) ContentsHandlerV1 {
	return ContentsHandlerV1{
		app: app,
	}
}

func (h ContentsHandlerV1) HandlePostContent(w http.ResponseWriter, r *http.Request) {
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

func (h ContentsHandlerV1) HandleViewContentDetails(w http.ResponseWriter, r *http.Request) {
	contentID := r.URL.Query().Get("content-id")

	content, err := h.app.GetContentDetails(r.Context(), contentID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response.WithJson(w, http.StatusOK, content)
}

func (h ContentsHandlerV1) HandleUpdateContent(w http.ResponseWriter, r *http.Request) {
	var cmdParams commands.UpdateContentDto
	err := json.NewDecoder(r.Body).Decode(&cmdParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cmdParams.UserID = r.Context().Value(middleware.ContextUserIDKey).(string)

	err = h.app.UpdateContent(r.Context(), &cmdParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h ContentsHandlerV1) HandleRemoveContent(w http.ResponseWriter, r *http.Request) {
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

func (h ContentsHandlerV1) HandleViewUsersContent(w http.ResponseWriter, r *http.Request) {
	requestedProfile := r.URL.Query().Get("user-id")

	if requestedProfile == "" {
		http.Error(w, "no profile specified", http.StatusBadRequest)
		return
	}

	contentList, err := h.app.GetContentOfUser(r.Context(), requestedProfile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response.WithJson(w, http.StatusOK, contentList)
}
