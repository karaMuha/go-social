package ports

import (
	"context"

	"github.com/karaMuha/go-social/posts/application/domain"
)

type ICommentsRepository interface {
	Create(ctx context.Context, comment *domain.Comment) error
	GetByPostID(ctx context.Context, commentID string) (*[]domain.Comment, error)
}
