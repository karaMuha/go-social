package ports

import (
	"context"

	"github.com/karaMuha/go-social/users/application/commands"
)

type IApplication interface {
	ICommands
	IQueries
}

type ICommands interface {
	RegisterUser(ctx context.Context, cmd commands.RegisterUserDto) error
}

type IQueries interface{}
