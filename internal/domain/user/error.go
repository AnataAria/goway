package user

type DomainError struct {
	message string
}

func (e *DomainError) Error() string {
	return e.message
}

func NewErrEmptyEmail() error {
	return &DomainError{message: "email cannot be empty"}
}

func NewErrInvalidEmail() error {
	return &DomainError{message: "email format is invalid"}
}

func NewErrEmptyPassword() error {
	return &DomainError{message: "password cannot be empty"}
}

func NewErrPasswordTooShort() error {
	return &DomainError{message: "password must be at least 6 characters"}
}

func NewErrEmptyName() error {
	return &DomainError{message: "name cannot be empty"}
}

func NewErrUserNotFound() error {
	return &DomainError{message: "user not found"}
}

func NewErrUserAlreadyExists() error {
	return &DomainError{message: "user already exists"}
}
