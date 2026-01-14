package http

import (
	"context"
	"errors"

	"github.com/AnataAria/goway/internal/domain/user"
	"github.com/AnataAria/goway/pkg/validator"
	"github.com/go-fuego/fuego"
	go_validator "github.com/go-playground/validator/v10"
)

func errorHandler(c context.Context, err error) error {
	_ = c

	var ve go_validator.ValidationErrors
	if errors.As(err, &ve) {
		return &fuego.HTTPError{
			Status: 400,
			Title:  "Validation Error",
			Err:    validator.ToErrorResponse(ve),
			Errors: toFuegoErrors(validator.ToErrorResponse(ve).Errors),
		}
	}

	var customVE *validator.ValidationError
	if errors.As(err, &customVE) {
		return &fuego.HTTPError{
			Status: 400,
			Title:  "Validation Error",
			Err:    customVE,
		}
	}

	var httpErr *fuego.HTTPError
	if errors.As(err, &httpErr) {
		return httpErr
	}

	switch {
	case errors.Is(err, user.ErrUserNotFound):
		return &fuego.HTTPError{Status: 404, Title: "Not Found", Detail: err.Error()}
	case errors.Is(err, user.ErrUserAlreadyExists):
		return &fuego.HTTPError{Status: 409, Title: "Conflict", Detail: err.Error()}
	case errors.Is(err, user.ErrInvalidEmail),
		errors.Is(err, user.ErrPasswordTooShort),
		errors.Is(err, user.ErrEmptyEmail),
		errors.Is(err, user.ErrEmptyPassword),
		errors.Is(err, user.ErrEmptyName):
		return &fuego.HTTPError{Status: 400, Title: "Bad Request", Detail: err.Error()}
	}

	return &fuego.HTTPError{
		Status: 500,
		Title:  "Internal Server Error",
		Detail: "An unexpected error occurred",
	}
}

func toFuegoErrors(errors []validator.FieldError) []fuego.ErrorItem {
	out := make([]fuego.ErrorItem, len(errors))
	for i, e := range errors {
		out[i] = fuego.ErrorItem{
			Name:   e.Field,
			Reason: e.Message,
		}
	}
	return out
}
