package auth

import (
	portAuth "github.com/AnataAria/goway/internal/port/in/auth"
)

type AuthAdapter struct {
	loginUseCase portAuth.LoginUseCase
}

func NewAuthAdapter(
	loginUseCase portAuth.LoginUseCase,
) *AuthAdapter {
	return &AuthAdapter{
		loginUseCase: loginUseCase,
	}
}
