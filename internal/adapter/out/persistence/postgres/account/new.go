package account

import (
	persistAccount "github.com/AnataAria/goway/internal/port/out/persistence/account"
	"github.com/jmoiron/sqlx"
)

func NewAccountRepository(db *sqlx.DB) persistAccount.AccountRepository {
	return NewPostgresAccountRepository(db)
}
