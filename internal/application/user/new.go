package user

import (
	persistUser "github.com/AnataAria/goway/internal/port/out/persistence/user"
)

func NewRegisterUseCase(userRepo persistUser.UserRepository) *RegisterHandler {
	return NewRegisterHandler(userRepo)
}

func NewGetUserUseCase(userRepo persistUser.UserRepository) *GetUserHandler {
	return NewGetUserHandler(userRepo)
}
