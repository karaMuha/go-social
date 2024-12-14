package driver

import (
	"context"

	"github.com/karaMuha/go-social/contents/application/commands"
	"github.com/karaMuha/go-social/contents/application/domain"
)

type IApplication interface {
	ICommands
	IQueries
}

type ICommands interface {
	PostContent(ctx context.Context, cmd *commands.PostContentDto) (string, error)
	UpdateContent(ctx context.Context, cmd *commands.UpdateContentDto) error
	RemoveContent(ctx context.Context, cmd *commands.RemoveContentDto) error
}

type IQueries interface {
	GetContentDetails(ctx context.Context, postID string) (*domain.Post, error)
}
