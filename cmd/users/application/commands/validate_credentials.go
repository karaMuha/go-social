package commands

import (
	"context"
	"fmt"

	"github.com/karaMuha/go-social/users/application/domain"
	"github.com/karaMuha/go-social/users/application/ports/driven"
)

type ValidateCredentialsDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ValidateUserCommand struct {
	usersRepository driven.IUsersRepsitory
}

func NewValidateUserCommand(usersRepository driven.IUsersRepsitory) ValidateUserCommand {
	return ValidateUserCommand{
		usersRepository: usersRepository,
	}
}

func (c ValidateUserCommand) ValidateUser(ctx context.Context, cmd *ValidateCredentialsDto) (string, error) {
	user, err := c.usersRepository.GetByEmail(ctx, cmd.Email)
	if err != nil {
		return "", fmt.Errorf("error logging in user: %v", err)
	}

	err = domain.ValidatePassword(user.Password, cmd.Password)
	if err != nil {
		return "", fmt.Errorf("error logging in user: %v", err)
	}

	return user.ID, nil
}
