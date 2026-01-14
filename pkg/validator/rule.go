package validator

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func passwordRule(fl validator.FieldLevel) bool {
	p := fl.Field().String()
	return strings.ContainsAny(p, "0123456789") &&
		strings.ContainsAny(p, "!@#$%^&*")
}
