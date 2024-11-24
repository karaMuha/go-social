package application

import ports "github.com/karaMuha/go-social/posts/application/ports/driver"

type Application struct {
	Commands
	Queries
}

type Commands struct{}

type Queries struct{}

var _ ports.IApplication = (*Application)(nil)

func New() Application {
	return Application{}
}
