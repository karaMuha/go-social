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
	app := application.New(usersRepository)

	// setup driver adapters
	usersHandlerV1 := rest.NewUsersHandlerV1(app)
	setupRoutes(mono.Mux(), usersHandlerV1)
	return nil
}

func setupRoutes(router *http.ServeMux, usersHandlerV1 rest.UsersHandlerV1) {
	v1 := http.NewServeMux()
	v1.HandleFunc("POST /users", usersHandlerV1.RegisterUserHandler)

	router.Handle("/v1/", http.StripPrefix("/v1", v1))
}
