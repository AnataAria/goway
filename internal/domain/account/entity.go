package account

type Account struct {
	ID        string
	Email     *Email
	Password  *Password
	Name      string
	CreatedAt int64
}

func NewAccount(id string, email *Email, password *Password, name string) *Account {
	return &Account{
		ID:        id,
		Email:     email,
		Password:  password,
		Name:      name,
		CreatedAt: 0,
	}
}

func (a *Account) GetID() string {
	return a.ID
}

func (a *Account) GetEmail() *Email {
	return a.Email
}

func (a *Account) GetPassword() *Password {
	return a.Password
}

func (a *Account) GetName() string {
	return a.Name
}

func (a *Account) ChangeName(name string) error {
	if name == "" {
		return NewErrEmptyName()
	}
	a.Name = name
	return nil
}

func (a *Account) ChangeEmail(email *Email) error {
	if email == nil {
		return NewErrInvalidEmail()
	}
	a.Email = email
	return nil
}
