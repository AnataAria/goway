package auth

import (
	"github.com/AnataAria/goway/internal/application/auth"
	portAuth "github.com/AnataAria/goway/internal/port/in/auth"
	persistAccount "github.com/AnataAria/goway/internal/port/out/persistence/account"
	security "github.com/AnataAria/goway/internal/port/out/security"
	"github.com/go-fuego/fuego"
)

type AuthHandler struct {
	accountRepo    persistAccount.AccountRepository
	passwordHasher security.PasswordHasher
	tokenGenerator security.TokenGenerator
	loginUseCase   *auth.LoginHandler
}

func NewAuthHandler(
	accountRepo persistAccount.AccountRepository,
	passwordHasher security.PasswordHasher,
	tokenGenerator security.TokenGenerator,
) *AuthHandler {
	return &AuthHandler{
		accountRepo:    accountRepo,
		passwordHasher: passwordHasher,
		tokenGenerator: tokenGenerator,
		loginUseCase:   auth.NewLoginUseCase(accountRepo, passwordHasher, tokenGenerator),
	}
}

func (h *AuthHandler) Login(c fuego.ContextWithBody[LoginRequest]) (LoginResponse, error) {
	body, err := c.Body()
	if err != nil {
		return LoginResponse{}, err
	}

	response, err := h.loginUseCase.Login(&portAuth.LoginRequest{
		Email:    body.Email,
		Password: body.Password,
	})
	if err != nil {
		return LoginResponse{}, err
	}

	return LoginResponse{
		Token: response.Token,
	}, nil
}
