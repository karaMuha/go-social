package posts

import (
	"context"
	"net/http"

	"github.com/karaMuha/go-social/contents/application"
	"github.com/karaMuha/go-social/contents/postgres"
	"github.com/karaMuha/go-social/contents/rest/v1"
	"github.com/karaMuha/go-social/internal/monolith"
)

type Module struct{}

func (m *Module) Startup(ctx context.Context, mono monolith.IMonolith) error {
	// setup driven adapters
	postsRepository := postgres.NewPostsRepository(mono.DB())

	// setup application
	app := application.New(postsRepository)

	// setup driver adapters
	postsHandlerV1 := rest.NewContentsHandlerV1(app)
	setupEndpoints(mono.Mux(), postsHandlerV1)

	return nil
}

func setupEndpoints(mux *http.ServeMux, postsHandlerV1 rest.PostsHandlerV1) {
	postsV1 := http.NewServeMux()
	postsV1.HandleFunc("POST /post-content", postsHandlerV1.HandlePostContent)
	postsV1.HandleFunc("GET /view-content-details", postsHandlerV1.HandleViewContentDetails)
	postsV1.HandleFunc("POST /update-content", postsHandlerV1.HandleUpdateContent)
	postsV1.HandleFunc("POST /remove-content", postsHandlerV1.HandleRemoveContent)

	mux.Handle("/v1/posts/", http.StripPrefix("/v1/posts", postsV1))
}
