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
	UpdatePost(ctx context.Context, cmd *commands.UpdatePostDto) error
	DeletePost(ctx context.Context, cmd *commands.DeletePostDto) error
}

type IQueries interface {
	GetPost(ctx context.Context, query *queries.GetPostDto) (*domain.Post, error)
}
