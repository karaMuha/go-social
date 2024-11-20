package ports

import (
	"context"

	"github.com/karaMuha/go-social/posts/application/domain"
)

type IPostRepository interface {
	GetByID(ctx context.Context, postID string) (*domain.Post, error)
	Create(ctx context.Context, post *domain.Post) error
	Delete(ctx context.Context, postID string) error
	Update(ctx context.Context, post *domain.Post) error
	GetUserFeed(ctx context.Context, userID string, pagination *domain.PaginatedFeedQuery) (*[]domain.PostWithMetadata, error)
}
