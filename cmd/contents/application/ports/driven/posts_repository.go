package driven

import (
	"context"

	"github.com/karaMuha/go-social/contents/application/domain"
)

type PostsRepository interface {
	CreateEntry(ctx context.Context, post *domain.Content) (string, error)
	GetByID(ctx context.Context, postID string) (*domain.Content, error)
	UpdateEntry(ctx context.Context, post *domain.Content) error
	DeleteEntry(ctx context.Context, postID string) error
}
