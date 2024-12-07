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
	followersRepository := postgres.NewFollowersRepository(mono.DB())

	// setup application
	app := application.New(
		usersRepository,
		followersRepository,
		mono.MailServer(),
		mono.Config().PrivateKeyPath,
	)

	// setup driver adapters
	usersHandlerV1 := rest.NewUsersHandlerV1(app, mono.TokenProvider())
	followersHandlerV1 := rest.NewFollowersHandlerV1(app)
	setupRoutes(mono.Mux(), usersHandlerV1, followersHandlerV1)

	return nil
}

func setupRoutes(
	router *http.ServeMux,
	usersHandlerV1 rest.UsersHandlerV1,
	followersHandlerV1 rest.FollowersHandlerV1,
) {
	usersV1 := http.NewServeMux()
	usersV1.HandleFunc("POST /", usersHandlerV1.SignupHandler)
	usersV1.HandleFunc("PUT /confirm", usersHandlerV1.ConfirmHandler)
	usersV1.HandleFunc("GET /email/{email}", usersHandlerV1.GetByEmailHandler)
	usersV1.HandleFunc("GET /{id}", usersHandlerV1.GetByIdHandler)
	usersV1.HandleFunc("POST /login", usersHandlerV1.LoginHandler)

	router.Handle("/v1/users/", http.StripPrefix("/v1/users", usersV1))

	followersV1 := http.NewServeMux()
	followersV1.HandleFunc("POST /", followersHandlerV1.FollowHandler)
	followersV1.HandleFunc("DELETE /", followersHandlerV1.UnfollowHandler)
	followersV1.HandleFunc("GET /{id}", followersHandlerV1.GetFollowersOfUser)

	router.Handle("/v1/followers/", http.StripPrefix("/v1/followers", followersV1))
}
