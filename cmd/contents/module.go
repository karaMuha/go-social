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
	contentsRepository := postgres.NewContentsRepository(mono.DB())

	// setup application
	app := application.New(contentsRepository)

	// setup driver adapters
	postsHandlerV1 := rest.NewContentsHandlerV1(app)
	setupEndpoints(mono.Mux(), postsHandlerV1)

	return nil
}

func setupEndpoints(mux *http.ServeMux, postsHandlerV1 rest.ContentsHandlerV1) {
	contentsV1 := http.NewServeMux()
	contentsV1.HandleFunc("POST /post-content", postsHandlerV1.HandlePostContent)
	contentsV1.HandleFunc("GET /view-content-details", postsHandlerV1.HandleViewContentDetails)
	contentsV1.HandleFunc("POST /update-content", postsHandlerV1.HandleUpdateContent)
	contentsV1.HandleFunc("POST /remove-content", postsHandlerV1.HandleRemoveContent)
	contentsV1.HandleFunc("GET /view-users-content", postsHandlerV1.HandleViewUsersContent)

	mux.Handle("/v1/posts/", http.StripPrefix("/v1/posts", contentsV1))
}
