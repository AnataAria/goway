package user

import (
	domainUser "github.com/AnataAria/goway/internal/domain/user"
)

func ToDomain(row UserRow) (*domainUser.User, error) {
	email, err := domainUser.NewEmail(row.Email)
	if err != nil {
		return nil, err
	}

	password, err := domainUser.NewPassword(row.Password)
	if err != nil {
		return nil, err
	}

	return domainUser.NewUser(row.ID, email, password, row.Name), nil
}

func ToRow(user *domainUser.User) UserRow {
	return UserRow{
		ID:       user.GetID(),
		Email:    user.GetEmail().Value(),
		Password: user.GetPassword().Value(),
		Name:     user.GetName(),
	}
}
