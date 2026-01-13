package auth

import (
	persistAccount "github.com/AnataAria/goway/internal/port/out/persistence/account"
	security "github.com/AnataAria/goway/internal/port/out/security"
)

type AuthUserCase struct {
	accountRepo    persistAccount.AccountRepository
	passwordHasher security.PasswordHasher
	tokenGenerator security.TokenGenerator
}

func New(
	accountRepo persistAccount.AccountRepository,
	passwordHasher security.PasswordHasher,
	tokenGenerator security.TokenGenerator,
) *AuthUserCase {
	return &AuthUserCase{
		accountRepo:    accountRepo,
		passwordHasher: passwordHasher,
		tokenGenerator: tokenGenerator,
	}
}
