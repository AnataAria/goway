package user

import (
	portUser "github.com/AnataAria/goway/internal/port/in/user"
)

type UserAdapter struct {
	userUseCase portUser.UserUseCase
}

func NewUserAdapter(
	userUseCase portUser.UserUseCase,
) *UserAdapter {
	return &UserAdapter{
		userUseCase: userUseCase,
	}
}
