package commands

import (
	"context"
	"fmt"

	"github.com/karaMuha/go-social/internal/mailer"
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
	mailer    mailer.Mailer
}

func NewSignupUserCommand(usersRepo ports.IUsersRepsitory, mailServer mailer.Mailer) SignupUserCommand {
	return SignupUserCommand{
		usersRepo: usersRepo,
		mailer:    mailServer,
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

	err = c.mailer.SendRegistrationMail(registration.Email, registration.RegistrationToken)
	if err != nil {
		return fmt.Errorf("error registering user: %w", err)
	}

	return nil
}
