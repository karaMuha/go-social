package driven

import (
	"context"

	"github.com/karaMuha/go-social/contents/application/domain"
)

type ContentsRepository interface {
	CreateEntry(ctx context.Context, content *domain.Content) (string, error)
	GetByID(ctx context.Context, contentID string) (*domain.Content, error)
	UpdateEntry(ctx context.Context, content *domain.Content) error
	DeleteEntry(ctx context.Context, contentID string) error
	GetAllOfUser(ctx context.Context, userID string) ([]*domain.Content, error)
}
