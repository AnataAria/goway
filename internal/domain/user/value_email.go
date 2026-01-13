package user

import (
	"regexp"
)

type Email struct {
	value string
}

func NewEmail(email string) (*Email, error) {
	if email == "" {
		return nil, NewErrEmptyEmail()
	}

	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	if !regexp.MustCompile(pattern).MatchString(email) {
		return nil, NewErrInvalidEmail()
	}

	return &Email{value: email}, nil
}

func (e *Email) Value() string {
	return e.value
}

func (e *Email) Equals(other *Email) bool {
	if other == nil {
		return false
	}
	return e.value == other.value
}

func (e *Email) String() string {
	return e.value
}
