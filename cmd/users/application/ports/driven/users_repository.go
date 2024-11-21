package ports

import (
	"context"

	"github.com/karaMuha/go-social/users/application/domain"
)

type IUsersRepsitory interface {
	GetByID(ctx context.Context, userID string) (*domain.User, error)
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
	Register(ctx context.Context, user *domain.User) (*domain.User, error)
	// Activate(ctx context.Context, token string) error
	Delete(ctx context.Context, userID string) error
}
