package commands

import (
	"context"
	"fmt"

	"github.com/karaMuha/go-social/users/application/domain"
	"github.com/karaMuha/go-social/users/application/ports/driven"
)

type LoginUserDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserCommand struct {
	usersRepository driven.IUsersRepsitory
}

func NewLoginUserCommand(usersRepository driven.IUsersRepsitory) LoginUserCommand {
	return LoginUserCommand{
		usersRepository: usersRepository,
	}
}

func (c LoginUserCommand) LoginUser(ctx context.Context, cmd *LoginUserDto) (string, error) {
	user, err := c.usersRepository.GetByEmail(ctx, cmd.Email)
	if err != nil {
		return "", fmt.Errorf("error logging in user: %v", err)
	}

	token, err := domain.Login(user.ID, user.Password, cmd.Password)
	if err != nil {
		return "", fmt.Errorf("error logging in user: %v", err)
	}

	return token, nil
}
