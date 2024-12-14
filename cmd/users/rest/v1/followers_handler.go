package rest

import (
	"encoding/json"
	"net/http"

	"github.com/karaMuha/go-social/internal/middleware"
	"github.com/karaMuha/go-social/users/application/commands"
	"github.com/karaMuha/go-social/users/application/ports/driver"
)

type FollowersHandlerV1 struct {
	app driver.IApplication
}

func NewFollowersHandlerV1(app driver.IApplication) FollowersHandlerV1 {
	return FollowersHandlerV1{
		app: app,
	}
}

func (h FollowersHandlerV1) FollowHandler(w http.ResponseWriter, r *http.Request) {
	var cmdParams commands.FollowUserDto
	err := json.NewDecoder(r.Body).Decode(&cmdParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cmdParams.UserID = r.Context().Value(middleware.ContextUserIDKey).(string)

	err = h.app.FollowUser(r.Context(), &cmdParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (h FollowersHandlerV1) UnfollowHandler(w http.ResponseWriter, r *http.Request) {
	var cmdParams commands.UnfollowUserDto
	err := json.NewDecoder(r.Body).Decode(&cmdParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cmdParams.UserID = r.Context().Value(middleware.ContextUserIDKey).(string)

	err = h.app.UnfollowUser(r.Context(), &cmdParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (h FollowersHandlerV1) ListFollowersOfUser(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user-id")
	followerList, err := h.app.GetFollowersOfUser(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	responseJson, err := json.Marshal(followerList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJson)
}
