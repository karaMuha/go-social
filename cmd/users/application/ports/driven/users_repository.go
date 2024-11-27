package driven

import (
	"context"

	"github.com/karaMuha/go-social/users/application/domain"
)

type IUsersRepsitory interface {
	GetByID(ctx context.Context, userID string) (*domain.Registration, error)
	GetByEmail(ctx context.Context, email string) (*domain.Registration, error)
	CreateEntry(ctx context.Context, user *domain.Registration) error
	ActivateUser(ctx context.Context, userID string) error
	DeleteEntry(ctx context.Context, userID string) error
}
