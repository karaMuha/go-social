package ports

type IApplication interface {
	ICommands
	IQueries
}

type ICommands interface{}

type IQueries interface{}
