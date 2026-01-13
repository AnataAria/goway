package account

type Password struct {
	value string
}

func NewPassword(password string) (*Password, error) {
	if password == "" {
		return nil, NewErrEmptyPassword()
	}

	if len(password) < 6 {
		return nil, NewErrPasswordTooShort()
	}

	return &Password{value: password}, nil
}

func (p *Password) Value() string {
	return p.value
}

func (p *Password) Equals(other *Password) bool {
	if other == nil {
		return false
	}
	return p.value == other.value
}
