package ports

import "context"

type Role struct {
	ID    string
	Name  string
	Level string
}

type IRolesRepository interface {
	GetByName(ctx context.Context, roleName string)
}
