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

func (h UsersHandlerV1) UserSignupHandler(w http.ResponseWriter, r *http.Request) {
	var cmdParams commands.RegisterUserDto
	err := json.NewDecoder(r.Body).Decode(&cmdParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.app.SignupUser(r.Context(), &cmdParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (h UsersHandlerV1) UserConfirmHandler(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	token := r.URL.Query().Get("token")
	cmdParams := commands.ConfirmUserDto{
		Email: email,
		Token: token,
	}

	err := h.app.ConfirmUser(r.Context(), &cmdParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
