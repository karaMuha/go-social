package commands

import (
	"context"
	"fmt"

	"github.com/karaMuha/go-social/users/application/domain"
	ports "github.com/karaMuha/go-social/users/application/ports/driven"
)

type RegisterUserDto struct {
	Email    string
	Username string
	Password string
}

type RegisterUserHandler struct {
	usersRepo ports.IUsersRepsitory
}

func NewRegisterUserHandler(usersRepo ports.IUsersRepsitory) RegisterUserHandler {
	return RegisterUserHandler{
		usersRepo: usersRepo,
	}
}

func (h RegisterUserHandler) RegisterUser(ctx context.Context, cmd RegisterUserDto) (*domain.User, error) {
	user, err := domain.RegisterUser(cmd.Username, cmd.Email, cmd.Password)
	if err != nil {
		return nil, fmt.Errorf("error registering user: %w", err)
	}

	user, err = h.usersRepo.Register(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("error registering user: %w", err)
	}

	return user, nil
}
