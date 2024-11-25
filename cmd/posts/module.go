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
	v1 := http.NewServeMux()
	v1.HandleFunc("POST /posts", postsHandlerV1.HandleCreatePost)
	v1.HandleFunc("GET /posts/{id}", postsHandlerV1.HandleGetPost)
	v1.HandleFunc("PUT /posts/{id}", postsHandlerV1.HandleUpdatePost)

	router.Handle("/v1/posts", http.StripPrefix("/v1", v1))
}
