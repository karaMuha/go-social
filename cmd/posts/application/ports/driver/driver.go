package driver

import (
	"context"

	"github.com/karaMuha/go-social/posts/application/commands"
	"github.com/karaMuha/go-social/posts/application/domain"
	"github.com/karaMuha/go-social/posts/application/queries"
)

type IApplication interface {
	ICommands
	IQueries
}

type ICommands interface {
	CreatePost(ctx context.Context, cmd *commands.CreatePostDto) error
}

type IQueries interface {
	GetPost(ctx context.Context, query *queries.GetPostDto) (*domain.Post, error)
}
