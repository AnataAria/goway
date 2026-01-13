package auth

import (
	domainAccount "github.com/AnataAria/goway/internal/domain/account"
	portAuth "github.com/AnataAria/goway/internal/port/in/auth"
)

func (h *AuthUserCase) Login(req *portAuth.LoginRequest) (*portAuth.LoginResponse, error) {
	email, err := domainAccount.NewEmail(req.Email)
	if err != nil {
		return nil, err
	}

	account, err := h.accountRepo.FindByEmail(email.Value())
	if err != nil {
		return nil, domainAccount.NewErrInvalidCredentials()
	}
	if account == nil {
		return nil, domainAccount.NewErrInvalidCredentials()
	}

	err = h.passwordHasher.Verify(account.GetPassword().Value(), req.Password)
	if err != nil {
		return nil, domainAccount.NewErrInvalidCredentials()
	}

	token, err := h.tokenGenerator.Generate(account.GetID())
	if err != nil {
		return nil, err
	}

	return &portAuth.LoginResponse{
		Token: token,
	}, nil
}
