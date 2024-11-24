package posts

import (
	"context"
	"net/http"

	"github.com/karaMuha/go-social/internal/monolith"
	"github.com/karaMuha/go-social/posts/application"
	"github.com/karaMuha/go-social/posts/rest/v1"
)

type Module struct{}

func (m *Module) Startup(ctx context.Context, mono monolith.IMonolith) error {
	// setup driven adapters

	// setup application
	app := application.New()

	// setup driver adapters
	postsHandlerV1 := rest.NewPostsHandlerV1(app)
	setupRoutes(mono.Mux(), postsHandlerV1)
	return nil
}

func setupRoutes(router *http.ServeMux, postsHandlerV1 rest.PostsHandlerV1) {
	v1 := http.NewServeMux()
	v1.HandleFunc("POST /posts", postsHandlerV1.PostCreationHandler)

	router.Handle("/v1/", http.StripPrefix("/v", v1))
}
