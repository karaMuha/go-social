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
	)

	// setup driver adapters
	usersHandlerV1 := rest.NewUsersHandlerV1(app, mono.TokenProvider())
	followersHandlerV1 := rest.NewFollowersHandlerV1(app)
	setupEndpoints(mono.Mux(), usersHandlerV1, followersHandlerV1)

	return nil
}

func setupEndpoints(
	mux *http.ServeMux,
	usersHandlerV1 rest.UsersHandlerV1,
	followersHandlerV1 rest.FollowersHandlerV1,
) {
	usersV1 := http.NewServeMux()
	usersV1.HandleFunc("POST /signup-for-registration", usersHandlerV1.HandleSignup)
	usersV1.HandleFunc("POST /confirm-registration", usersHandlerV1.HandleConfirm)
	usersV1.HandleFunc("GET /view-user-details", usersHandlerV1.HandleViewUserDetails)
	usersV1.HandleFunc("POST /login", usersHandlerV1.HandleLogin)

	mux.Handle("/v1/users/", http.StripPrefix("/v1/users", usersV1))

	followersV1 := http.NewServeMux()
	followersV1.HandleFunc("POST /follow-user", followersHandlerV1.FollowHandler)
	followersV1.HandleFunc("POST /unfollow-user", followersHandlerV1.UnfollowHandler)
	followersV1.HandleFunc("GET /list-followers-of-user", followersHandlerV1.ListFollowersOfUser)

	mux.Handle("/v1/followers/", http.StripPrefix("/v1/followers", followersV1))
}
