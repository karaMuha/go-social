package application

import (
	"github.com/karaMuha/go-social/users/application/commands"
	"github.com/karaMuha/go-social/users/application/domain"
	drivenPorts "github.com/karaMuha/go-social/users/application/ports/driven"
	driverPorts "github.com/karaMuha/go-social/users/application/ports/driver"
)

type Application struct {
	appCommands
	appQueries
}

type appCommands struct {
	commands.SignupUserCommand
}

type appQueries struct{}

var _ driverPorts.IApplication = (*Application)(nil)

func New(usersRepo drivenPorts.IUsersRepsitory) Application {
	domain.InitValidator()
	return Application{
		appCommands: appCommands{
			SignupUserCommand: commands.NewSignupUserCommand(usersRepo),
		},
	}
}
