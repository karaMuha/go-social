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
}

var _ driver.IApplication = (*Application)(nil)

func New(postsRepository driven.PostsRepository) Application {
	domain.InitValidator()
	return Application{
		appCommands: appCommands{
			PostContentCommand:   commands.NewPostContentCommand(postsRepository),
			UpdateContentCommand: commands.NewUpdateContentCommand(postsRepository),
			RemoveContentCommand: commands.NewRemoveContentCommand(postsRepository),
		},
		appQueries: appQueries{
			GetContentDetailsQuery: queries.NewGetContentDetailsQuery(postsRepository),
		},
	}
}
