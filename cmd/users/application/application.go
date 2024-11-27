package application

import (
	"github.com/karaMuha/go-social/internal/mailer"
	"github.com/karaMuha/go-social/users/application/commands"
	"github.com/karaMuha/go-social/users/application/domain"
	"github.com/karaMuha/go-social/users/application/ports/driven"
	"github.com/karaMuha/go-social/users/application/ports/driver"
	"github.com/karaMuha/go-social/users/application/queries"
)

type Application struct {
	appCommands
	appQueries
}

type appCommands struct {
	commands.SignupUserCommand
	commands.ConfirmUserCommand
}

type appQueries struct {
	queries.GetUserByEmailQuery
}

var _ driver.IApplication = (*Application)(nil)

func New(usersRepo driven.IUsersRepsitory, mailServer mailer.Mailer) Application {
	domain.InitValidator()
	return Application{
		appCommands: appCommands{
			SignupUserCommand:  commands.NewSignupUserCommand(usersRepo, mailServer),
			ConfirmUserCommand: commands.NewConfirmUserCommand(usersRepo),
		},
		appQueries: appQueries{
			GetUserByEmailQuery: queries.NewGetUserByEmailQuery(usersRepo),
		},
	}
}
