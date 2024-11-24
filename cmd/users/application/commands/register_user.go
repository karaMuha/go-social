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

type RegisterUserCommand struct {
	usersRepo ports.IUsersRepsitory
}

func NewRegisterUserCommand(usersRepo ports.IUsersRepsitory) RegisterUserCommand {
	return RegisterUserCommand{
		usersRepo: usersRepo,
	}
}

// Check within the domain logic if email and username already exist.
// Currently these checks depend on the postgres implementation.
// Take case insensivity for email (and username?) into account.
func (c RegisterUserCommand) RegisterUser(ctx context.Context, cmd RegisterUserDto) error {
	registration, err := domain.RegisterUser(cmd.Username, cmd.Email, cmd.Password)
	if err != nil {
		return fmt.Errorf("error registering user: %w", err)
	}

	err = c.usersRepo.CreateEntry(ctx, registration)
	if err != nil {
		return fmt.Errorf("error registering user: %w", err)
	}

	return nil
}
