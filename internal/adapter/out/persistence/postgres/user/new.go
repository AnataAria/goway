package user

import (
	persistUser "github.com/AnataAria/goway/internal/port/out/persistence/user"
	"github.com/jmoiron/sqlx"
)

func NewUserRepository(db *sqlx.DB) persistUser.UserRepository {
	return NewPostgresUserRepository(db)
}
