package ports

import (
	"context"
	"time"

	"github.com/karaMuha/go-social/users/application/domain"
)

type IUsersRepsitory interface {
	GetByID(ctx context.Context, userID string) (*domain.User, error)
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
	Register(ctx context.Context, user *domain.User, token string, exp time.Duration) error
	Activate(ctx context.Context, token string) error
	Delete(ctx context.Context, userID string) error
}
