package driven

import (
	"context"

	"github.com/karaMuha/go-social/posts/application/domain"
)

type PostsRepository interface {
	CreateEntry(ctx context.Context, post *domain.Post) (string, error)
	GetByID(ctx context.Context, postID string) (*domain.Post, error)
	UpdateEntry(ctx context.Context, post *domain.Post) error
	DeleteEntry(ctx context.Context, postID string) error
}
