package ports

type IValidator interface {
	Check(v any) error
}
