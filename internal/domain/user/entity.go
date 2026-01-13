package user

type User struct {
	ID        string
	Email     *Email
	Password  *Password
	Name      string
	CreatedAt int64
}

func NewUser(id string, email *Email, password *Password, name string) *User {
	return &User{
		ID:        id,
		Email:     email,
		Password:  password,
		Name:      name,
		CreatedAt: 0,
	}
}

func (u *User) GetID() string {
	return u.ID
}

func (u *User) GetEmail() *Email {
	return u.Email
}

func (u *User) GetPassword() *Password {
	return u.Password
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) ChangeName(name string) error {
	if name == "" {
		return NewErrEmptyName()
	}
	u.Name = name
	return nil
}

func (u *User) ChangeEmail(email *Email) error {
	if email == nil {
		return NewErrInvalidEmail()
	}
	u.Email = email
	return nil
}
