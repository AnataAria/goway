package validator

import (
	"github.com/ettle/strcase"
	"github.com/go-playground/validator/v10"
)

func (g *CustomValidator) Validate(i any) *ValidationError {
	if err := g.v.Struct(i); err != nil {
		ve := err.(validator.ValidationErrors)
		return ToErrorResponse(ve)
	}
	return nil
}

func ToErrorResponse(ve validator.ValidationErrors) *ValidationError {
	errors := make([]FieldError, 0)
	for _, e := range ve {
		errors = append(errors, FieldError{
			Field:   strcase.ToSnake(e.Field()),
			Message: messageFor(e),
		})
	}

	return &ValidationError{&ErrorResponse{
		Code:    "VALIDATION_ERROR",
		Message: "Invalid request",
		Errors:  errors,
	}}
}

func messageFor(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return "is required"
	case "email":
		return "must be a valid email"
	case "min":
		return "must be at least " + e.Param() + " characters"
	default:
		return "is invalid"
	}
}
