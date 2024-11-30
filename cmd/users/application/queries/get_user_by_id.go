package queries

import (
	"context"

	"github.com/karaMuha/go-social/users/application/domain"
	"github.com/karaMuha/go-social/users/application/ports/driven"
)

type GetUserByIdQuery struct {
	usersRepository driven.IUsersRepsitory
}

func NewGetUserByIdQuery(usersRepository driven.IUsersRepsitory) GetUserByIdQuery {
	return GetUserByIdQuery{
		usersRepository: usersRepository,
	}
}

func (q GetUserByEmailQuery) GetUserByID(ctx context.Context, userID string) (*domain.Registration, error) {
	return q.usersRepository.GetByID(ctx, userID)
}
