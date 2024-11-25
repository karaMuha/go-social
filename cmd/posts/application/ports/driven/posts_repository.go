package driven

import (
	"context"

	"github.com/karaMuha/go-social/posts/application/domain"
)

type PostsRepository interface {
	CreateEntry(ctx context.Context, post *domain.Post) error
	GetByID(ctx context.Context, postID string) (*domain.Post, error)
}
