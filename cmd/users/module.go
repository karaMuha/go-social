package users

import (
	"context"
	"net/http"

	"github.com/karaMuha/go-social/internal/monolith"
	"github.com/karaMuha/go-social/users/application"
	"github.com/karaMuha/go-social/users/postgres"
	"github.com/karaMuha/go-social/users/rest/v1"
)

type Module struct{}

func (m *Module) Startup(ctx context.Context, mono monolith.IMonolith) error {
	// setup driven adapters
	usersRepository := postgres.NewUsersRepository(mono.DB())

	// setup application
	app := application.New(usersRepository, mono.MailServer())

	// setup driver adapters
	usersHandlerV1 := rest.NewUsersHandlerV1(app)
	setupRoutes(mono.Mux(), usersHandlerV1)

	return nil
}

func setupRoutes(router *http.ServeMux, usersHandlerV1 rest.UsersHandlerV1) {
	usersV1 := http.NewServeMux()
	usersV1.HandleFunc("POST /", usersHandlerV1.UserSignupHandler)
	usersV1.HandleFunc("PUT /confirm", usersHandlerV1.UserConfirmHandler)
	usersV1.HandleFunc("GET /{email}", usersHandlerV1.UserGetByEmailHandler)

	router.Handle("/v1/users/", http.StripPrefix("/v1/users", usersV1))
}
