package account

import (
	domainAccount "github.com/AnataAria/goway/internal/domain/account"
)

func ToDomain(row AccountRow) (*domainAccount.Account, error) {
	email, err := domainAccount.NewEmail(row.Email)
	if err != nil {
		return nil, err
	}

	password, err := domainAccount.NewPassword(row.Password)
	if err != nil {
		return nil, err
	}

	return domainAccount.NewAccount(row.ID, email, password, row.Name), nil
}

func ToRow(account *domainAccount.Account) AccountRow {
	return AccountRow{
		ID:       account.GetID(),
		Email:    account.GetEmail().Value(),
		Password: account.GetPassword().Value(),
		Name:     account.GetName(),
	}
}
