package user

import (
	persistUser "github.com/AnataAria/goway/internal/port/out/persistence/user"
)

func NewUserAdapter(userRepo persistUser.UserRepository) *UserHandler {
	return NewUserHandler(userRepo)
}
