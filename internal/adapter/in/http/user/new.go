package user

import (
	"github.com/AnataAria/goway/internal/application/user"
	persistUser "github.com/AnataAria/goway/internal/port/out/persistence/user"
)

type UserAdapter struct {
	userRepo        persistUser.UserRepository
	registerUseCase *user.RegisterHandler
	getUserUseCase  *user.GetUserHandler
}

func NewUserAdapter(userRepo persistUser.UserRepository) *UserAdapter {
	return &UserAdapter{
		userRepo:        userRepo,
		registerUseCase: user.NewRegisterHandler(userRepo),
		getUserUseCase:  user.NewGetUserHandler(userRepo),
	}
}
