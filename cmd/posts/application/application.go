package application

import (
	"github.com/karaMuha/go-social/posts/application/commands"
	"github.com/karaMuha/go-social/posts/application/domain"
	"github.com/karaMuha/go-social/posts/application/ports/driven"
	"github.com/karaMuha/go-social/posts/application/ports/driver"
)

type Application struct {
	appCommands
	appQueries
}

type appCommands struct {
	commands.CreatePostCommand
}

type appQueries struct{}

var _ driver.IApplication = (*Application)(nil)

func New(postsRepository driven.PostsRepository) Application {
	domain.InitValidator()
	return Application{
		appCommands: appCommands{
			CreatePostCommand: commands.NewCreatePostCommand(postsRepository),
		},
		appQueries: appQueries{},
	}
}
