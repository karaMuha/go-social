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
	commands.FollowUserCommand
	commands.UnfollowUserCommand
}

type appQueries struct {
	queries.GetUserByEmailQuery
	queries.GetFollowerQuery
}

var _ driver.IApplication = (*Application)(nil)

func New(
	usersRepo driven.IUsersRepsitory,
	followersRepository driven.IFollowersRepository,
	mailServer mailer.Mailer,
	privateKeyPath string,
) Application {
	domain.InitValidator()
	domain.InitPrivateKey(privateKeyPath)
	return Application{
		appCommands: appCommands{
			SignupUserCommand:   commands.NewSignupUserCommand(usersRepo, mailServer),
			ConfirmUserCommand:  commands.NewConfirmUserCommand(usersRepo),
			FollowUserCommand:   commands.NewFollowUserCommand(followersRepository),
			UnfollowUserCommand: commands.NewUnfollowUserCommand(followersRepository),
		},
		appQueries: appQueries{
			GetUserByEmailQuery: queries.NewGetUserByEmailQuery(usersRepo),
			GetFollowerQuery:    queries.NewGetFollowerQuery(followersRepository),
		},
	}
}
