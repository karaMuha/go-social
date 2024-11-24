package application

import (
	"github.com/karaMuha/go-social/users/application/commands"
	"github.com/karaMuha/go-social/users/application/domain"
	"github.com/karaMuha/go-social/users/application/ports/driven"
	"github.com/karaMuha/go-social/users/application/ports/driver"
)

type Application struct {
	appCommands
	appQueries
}

type appCommands struct {
	commands.SignupUserCommand
}

type appQueries struct{}

var _ driver.IApplication = (*Application)(nil)

func New(usersRepo driven.IUsersRepsitory) Application {
	domain.InitValidator()
	return Application{
		appCommands: appCommands{
			SignupUserCommand: commands.NewSignupUserCommand(usersRepo),
		},
		appQueries: appQueries{},
	}
}
