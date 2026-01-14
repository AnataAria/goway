package validator

import "github.com/go-playground/validator/v10"

type CustomValidator struct {
	v *validator.Validate
}

func New() *CustomValidator {
	v := validator.New()

	v.RegisterValidation("password", passwordRule)

	return &CustomValidator{v: v}
}

func (cv *CustomValidator) Validator() *validator.Validate {
	return cv.v
}
