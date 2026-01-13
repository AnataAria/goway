package auth

import (
	"github.com/AnataAria/goway/internal/application/auth"
	persistAccount "github.com/AnataAria/goway/internal/port/out/persistence/account"
	security "github.com/AnataAria/goway/internal/port/out/security"
)

type AuthAdapter struct {
	accountRepo    persistAccount.AccountRepository
	passwordHasher security.PasswordHasher
	tokenGenerator security.TokenGenerator
	loginUseCase   *auth.LoginHandler
}

func NewAuthAdapter(
	accountRepo persistAccount.AccountRepository,
	passwordHasher security.PasswordHasher,
	tokenGenerator security.TokenGenerator,
) *AuthAdapter {
	return &AuthAdapter{
		accountRepo:    accountRepo,
		passwordHasher: passwordHasher,
		tokenGenerator: tokenGenerator,
		loginUseCase:   auth.NewLoginHandler(accountRepo, passwordHasher, tokenGenerator),
	}
}
