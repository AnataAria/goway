package account

import "github.com/AnataAria/goway/internal/domain/account"

type AccountRepository interface {
	Save(account *account.Account) error
	FindByID(id string) (*account.Account, error)
	FindByEmail(email string) (*account.Account, error)
	ExistsByEmail(email string) (bool, error)
}
