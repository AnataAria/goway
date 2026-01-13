package auth

import (
	persistAccount "github.com/AnataAria/goway/internal/port/out/persistence/account"
	security "github.com/AnataAria/goway/internal/port/out/security"
)

func NewAuthAdapter(
	accountRepo persistAccount.AccountRepository,
	passwordHasher security.PasswordHasher,
	tokenGenerator security.TokenGenerator,
) *AuthHandler {
	return NewAuthHandler(accountRepo, passwordHasher, tokenGenerator)
}
