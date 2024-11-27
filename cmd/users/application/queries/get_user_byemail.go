package queries

import (
	"context"

	"github.com/karaMuha/go-social/users/application/domain"
	"github.com/karaMuha/go-social/users/application/ports/driven"
)

type GetUserByEmailQuery struct {
	usersRepository driven.IUsersRepsitory
}

func NewGetUserByEmailQuery(usersRepository driven.IUsersRepsitory) GetUserByEmailQuery {
	return GetUserByEmailQuery{
		usersRepository: usersRepository,
	}
}

func (q GetUserByEmailQuery) GetUserByEmail(ctx context.Context, email string) (*domain.Registration, error) {
	return q.usersRepository.GetByEmail(ctx, email)
}
