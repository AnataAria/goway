package user

import (
	persistUser "github.com/AnataAria/goway/internal/port/out/persistence/user"
)

type UserUseCase struct {
	userRepo persistUser.UserRepository
}

func New(userRepo persistUser.UserRepository) *UserUseCase {
	return &UserUseCase{
		userRepo: userRepo,
	}
}
