package commands

import (
	"context"
	"fmt"
	"log"

	"github.com/karaMuha/go-social/internal/mailer"
	"github.com/karaMuha/go-social/users/application/domain"
	ports "github.com/karaMuha/go-social/users/application/ports/driven"
)

type SignupUserDto struct {
	Email    string
	Username string
	Password string
}

type SignupUserCommand struct {
	usersRepo ports.IUsersRepsitory
	mailer    mailer.IMailer
}

func NewSignupUserCommand(usersRepo ports.IUsersRepsitory, mailServer mailer.IMailer) SignupUserCommand {
	return SignupUserCommand{
		usersRepo: usersRepo,
		mailer:    mailServer,
	}
}

func (c SignupUserCommand) SignupUser(ctx context.Context, cmd *SignupUserDto) error {
	registration, err := domain.Signup(cmd.Username, cmd.Email, cmd.Password)
	if err != nil {
		return fmt.Errorf("error signing up user: %w", err)
	}

	userID, err := c.usersRepo.CreateEntry(ctx, registration)
	if err != nil {
		return fmt.Errorf("error signing up user: %w", err)
	}

	err = c.mailer.SendRegistrationMail(registration.Email, registration.RegistrationToken)
	if err != nil {
		// transaction with rollback?
		err = c.usersRepo.DeleteEntry(ctx, userID)
		if err != nil {
			log.Printf("Error deleting user entry after failed mail attempt: %v\n", err)
		}
		return fmt.Errorf("error signing up user: %w", err)
	}

	return nil
}
