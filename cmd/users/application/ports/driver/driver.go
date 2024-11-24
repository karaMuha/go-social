package driver

import (
	"context"

	"github.com/karaMuha/go-social/users/application/commands"
)

type IApplication interface {
	ICommands
	IQueries
}

type ICommands interface {
	SignupUser(ctx context.Context, cmd commands.RegisterUserDto) error
}

type IQueries interface{}
