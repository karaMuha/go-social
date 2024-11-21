package utils

import (
	"github.com/go-playground/validator/v10"
	ports "github.com/karaMuha/go-social/users/application/ports/utils"
)

var Validator Checker

type Checker struct {
	*validator.Validate
}

var _ ports.IValidator = (*Checker)(nil)

func InitValidator() {
	Validator = Checker{
		Validate: validator.New(),
	}
}

func (c Checker) Check(v any) error {
	return c.Struct(v)
}
