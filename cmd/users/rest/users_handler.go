package rest

import (
	"encoding/json"
	"net/http"

	"github.com/karaMuha/go-social/users/application"
	"github.com/karaMuha/go-social/users/application/commands"
)

type UsersHandler struct {
	app application.Application
}

func NewUsersHandler(app application.Application) UsersHandler {
	return UsersHandler{
		app: app,
	}
}

func (h UsersHandler) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody commands.RegisterUserDto
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.app.RegisterUser(r.Context(), requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
