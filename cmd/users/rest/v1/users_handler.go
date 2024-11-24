package rest

import (
	"encoding/json"
	"net/http"

	"github.com/karaMuha/go-social/users/application/commands"
	ports "github.com/karaMuha/go-social/users/application/ports/driver"
)

type UsersHandlerV1 struct {
	app ports.IApplication
}

func NewUsersHandlerV1(app ports.IApplication) UsersHandlerV1 {
	return UsersHandlerV1{
		app: app,
	}
}

func (h UsersHandlerV1) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
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
