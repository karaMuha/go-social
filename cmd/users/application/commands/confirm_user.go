package commands

import (
	"context"
	"fmt"

	"github.com/karaMuha/go-social/users/application/domain"
	"github.com/karaMuha/go-social/users/application/ports/driven"
)

type ConfirmUserDto struct {
	Email string
	Token string
}

type ConfirmUserCommand struct {
	usersRepository driven.IUsersRepsitory
}

func NewConfirmUserCommand(usersRepository driven.IUsersRepsitory) ConfirmUserCommand {
	return ConfirmUserCommand{
		usersRepository: usersRepository,
	}
}

func (c ConfirmUserCommand) ConfirmUser(ctx context.Context, cmd *ConfirmUserDto) error {
	user, err := c.usersRepository.GetByEmail(ctx, cmd.Email)
	if err != nil {
		return fmt.Errorf("error activating user: %v", err)
	}

	err = domain.Activate(user, cmd.Token)
	if err != nil {
		return fmt.Errorf("error activating user: %v", err)
	}

	err = c.usersRepository.ActivateUser(ctx, user.ID)
	if err != nil {
		return fmt.Errorf("error activating user: %v", err)
	}

	return nil
}
