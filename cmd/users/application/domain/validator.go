package domain

import "github.com/go-playground/validator/v10"

// Validator is used to validate incoming data.
// Data is validated according to the struct tags described in the domain model.
// Variable must be initialized on startup.
var validate *validator.Validate

func InitValidator() {
	validate = validator.New()
}
