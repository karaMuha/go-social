package posts

import (
	"context"
	"net/http"

	"github.com/karaMuha/go-social/internal/monolith"
	"github.com/karaMuha/go-social/posts/application"
	"github.com/karaMuha/go-social/posts/postgres"
	"github.com/karaMuha/go-social/posts/rest/v1"
)

type Module struct{}

func (m *Module) Startup(ctx context.Context, mono monolith.IMonolith) error {
	// setup driven adapters
	postsRepository := postgres.NewPostsRepository(mono.DB())

	// setup application
	app := application.New(postsRepository)

	// setup driver adapters
	postsHandlerV1 := rest.NewPostsHandlerV1(app)
	setupRoutes(mono.Mux(), postsHandlerV1)

	return nil
}

func setupRoutes(router *http.ServeMux, postsHandlerV1 rest.PostsHandlerV1) {
	postsV1 := http.NewServeMux()
	postsV1.HandleFunc("POST /", postsHandlerV1.HandleCreatePost)
	postsV1.HandleFunc("GET /{id}", postsHandlerV1.HandleGetPost)
	postsV1.HandleFunc("PUT /{id}", postsHandlerV1.HandleUpdatePost)
	postsV1.HandleFunc("DELETE /{id}", postsHandlerV1.HandleDeletePost)

	router.Handle("/v1/posts/", http.StripPrefix("/v1/posts", postsV1))
}
