package driver

import (
	"context"

	"github.com/karaMuha/go-social/posts/application/commands"
)

type IApplication interface {
	ICommands
	IQueries
}

type ICommands interface {
	CreatePost(ctx context.Context, post *commands.CreatePostDto) error
}

type IQueries interface{}
