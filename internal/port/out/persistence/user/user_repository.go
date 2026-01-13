package user

import "github.com/AnataAria/goway/internal/domain/user"

type UserRepository interface {
	Save(user *user.User) error
	FindByID(id string) (*user.User, error)
	FindByEmail(email string) (*user.User, error)
	ExistsByEmail(email string) (bool, error)
}
