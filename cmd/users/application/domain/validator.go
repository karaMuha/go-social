package domain

import "github.com/go-playground/validator/v10"

// validator is used to validate incoming data
// according to the struct tags described in the domain model
// and must be initialized on startup
var validate *validator.Validate

func InitValidator() {
	validate = validator.New()
}
