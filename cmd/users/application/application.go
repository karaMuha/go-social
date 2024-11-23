package application

import (
	"github.com/karaMuha/go-social/users/application/domain"
	ports "github.com/karaMuha/go-social/users/application/ports/driver"
)

type Application struct {
	Commands
	Queries
}

type Commands struct{}

type Queries struct{}

var _ ports.IApplication = (*Application)(nil)

func New() *Application {
	domain.InitValidator()
	return &Application{}
}
