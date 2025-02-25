package rest

import (
	"encoding/json"
	"net/http"
	"time"

	authtoken "github.com/karaMuha/go-social/internal/auth_token"
	"github.com/karaMuha/go-social/internal/http/response"
	"github.com/karaMuha/go-social/users/application/commands"
	"github.com/karaMuha/go-social/users/application/domain"
	ports "github.com/karaMuha/go-social/users/application/ports/driver"
)

type UsersHandlerV1 struct {
	app           ports.IApplication
	tokenProvider authtoken.ITokenProvider
}

func NewUsersHandlerV1(app ports.IApplication, tokenProvider authtoken.ITokenProvider) UsersHandlerV1 {
	return UsersHandlerV1{
		app:           app,
		tokenProvider: tokenProvider,
	}
}

func (h UsersHandlerV1) HandleSignup(w http.ResponseWriter, r *http.Request) {
	var cmdParams commands.SignupUserDto
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

func (h UsersHandlerV1) HandleConfirm(w http.ResponseWriter, r *http.Request) {
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

func (h UsersHandlerV1) HandleViewUserDetails(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	userID := r.URL.Query().Get("user-id")

	if email == "" && userID == "" {
		http.Error(w, "no query parameter provided", http.StatusBadRequest)
		return
	}

	if email != "" && userID != "" {
		http.Error(w, "too many query parameter provided", http.StatusBadRequest)
		return
	}

	var user *domain.Registration
	var err error

	if email != "" {
		user, err = h.app.GetUserByEmail(r.Context(), email)
	}

	if userID != "" {
		user, err = h.app.GetUserByID(r.Context(), userID)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response.WithJson(w, http.StatusOK, user)
}

func (h UsersHandlerV1) HandleLogin(w http.ResponseWriter, r *http.Request) {
	var cmdParams commands.ValidateCredentialsDto
	err := json.NewDecoder(r.Body).Decode(&cmdParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID, err := h.app.ValidateUser(r.Context(), &cmdParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	accessToken, err := h.tokenProvider.GenerateToken(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Secure:   true,
		HttpOnly: true,
		Expires:  time.Now().Add(time.Hour),
		Path:     "/",
	})
}
