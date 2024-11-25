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

type SignupUserCommand struct {
	usersRepo ports.IUsersRepsitory
}

func NewSignupUserCommand(usersRepo ports.IUsersRepsitory) SignupUserCommand {
	return SignupUserCommand{
		usersRepo: usersRepo,
	}
}

func (c SignupUserCommand) SignupUser(ctx context.Context, cmd RegisterUserDto) error {
	registration, err := domain.Signup(cmd.Username, cmd.Email, cmd.Password)
	if err != nil {
		return fmt.Errorf("error registering user: %w", err)
	}

	err = c.usersRepo.CreateEntry(ctx, registration)
	if err != nil {
		return fmt.Errorf("error registering user: %w", err)
	}

	return nil
}
