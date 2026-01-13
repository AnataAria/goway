package auth

import (
	persistAccount "github.com/AnataAria/goway/internal/port/out/persistence/account"
	security "github.com/AnataAria/goway/internal/port/out/security"
)

func NewLoginUseCase(
	accountRepo persistAccount.AccountRepository,
	passwordHasher security.PasswordHasher,
	tokenGenerator security.TokenGenerator,
) *LoginHandler {
	return NewLoginHandler(accountRepo, passwordHasher, tokenGenerator)
}
