package application

import (
	"github.com/karaMuha/go-social/contents/application/commands"
	"github.com/karaMuha/go-social/contents/application/domain"
	"github.com/karaMuha/go-social/contents/application/ports/driven"
	"github.com/karaMuha/go-social/contents/application/ports/driver"
	"github.com/karaMuha/go-social/contents/application/queries"
)

type Application struct {
	appCommands
	appQueries
}

type appCommands struct {
	commands.PostContentCommand
	commands.UpdateContentCommand
	commands.RemoveContentCommand
}

type appQueries struct {
	queries.GetContentDetailsQuery
	queries.GetContentOfUserQuery
}

var _ driver.IApplication = (*Application)(nil)

func New(contentsRepository driven.ContentsRepository) Application {
	domain.InitValidator()
	return Application{
		appCommands: appCommands{
			PostContentCommand:   commands.NewPostContentCommand(contentsRepository),
			UpdateContentCommand: commands.NewUpdateContentCommand(contentsRepository),
			RemoveContentCommand: commands.NewRemoveContentCommand(contentsRepository),
		},
		appQueries: appQueries{
			GetContentDetailsQuery: queries.NewGetContentDetailsQuery(contentsRepository),
			GetContentOfUserQuery:  queries.NewGetContentOfUserQuery(contentsRepository),
		},
	}
}
