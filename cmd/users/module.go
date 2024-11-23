package users

import (
	"context"
	"net/http"

	"github.com/karaMuha/go-social/internal/monolith"
	"github.com/karaMuha/go-social/users/application"
	"github.com/karaMuha/go-social/users/postgres"
	"github.com/karaMuha/go-social/users/rest"
)

type Module struct{}

func (m *Module) Startup(ctx context.Context, mono monolith.IMonolith) error {
	// setup driven adapters
	usersRepository := postgres.NewUsersRepository(mono.DB())

	// setup application
	app := application.New(usersRepository)

	// setup driver adapters
	usersHandler := rest.NewUsersHandler(app)
	setupRoutes(mono.Mux(), usersHandler)
	return nil
}

func setupRoutes(router *http.ServeMux, usersHandler rest.UsersHandler) {
	router.HandleFunc("POST /users", usersHandler.RegisterUserHandler)

}
