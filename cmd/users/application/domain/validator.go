package domain

import "github.com/go-playground/validator/v10"

// validator is used to validate incoming data
// according to the struct tags described in the domain model
// and must be initialized on startup
var val *validator.Validate

func InitValidator() {
	val = validator.New()
}
