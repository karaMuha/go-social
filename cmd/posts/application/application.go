package application

import (
	"github.com/karaMuha/go-social/posts/application/commands"
	"github.com/karaMuha/go-social/posts/application/domain"
	"github.com/karaMuha/go-social/posts/application/ports/driven"
	"github.com/karaMuha/go-social/posts/application/ports/driver"
	"github.com/karaMuha/go-social/posts/application/queries"
)

type Application struct {
	appCommands
	appQueries
}

type appCommands struct {
	commands.CreatePostCommand
	commands.UpdatePostCommand
}

type appQueries struct {
	queries.GetPostQuery
}

var _ driver.IApplication = (*Application)(nil)

func New(postsRepository driven.PostsRepository) Application {
	domain.InitValidator()
	return Application{
		appCommands: appCommands{
			CreatePostCommand: commands.NewCreatePostCommand(postsRepository),
			UpdatePostCommand: commands.NewUpdatePostCommand(postsRepository),
		},
		appQueries: appQueries{
			GetPostQuery: queries.NewGetPostQuery(postsRepository),
		},
	}
}
