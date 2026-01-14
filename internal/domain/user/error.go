package user

import "errors"

var (
	ErrEmptyEmail        = errors.New("email cannot be empty")
	ErrInvalidEmail      = errors.New("email format is invalid")
	ErrEmptyPassword     = errors.New("password cannot be empty")
	ErrPasswordTooShort  = errors.New("password must be at least 6 characters")
	ErrEmptyName         = errors.New("name cannot be empty")
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
)

func NewErrEmptyEmail() error        { return ErrEmptyEmail }
func NewErrInvalidEmail() error      { return ErrInvalidEmail }
func NewErrEmptyPassword() error     { return ErrEmptyPassword }
func NewErrPasswordTooShort() error  { return ErrPasswordTooShort }
func NewErrEmptyName() error         { return ErrEmptyName }
func NewErrUserNotFound() error      { return ErrUserNotFound }
func NewErrUserAlreadyExists() error { return ErrUserAlreadyExists }
