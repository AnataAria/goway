package auth

import (
	"github.com/go-fuego/fuego"
)

func (h *AuthAdapter) Login(c fuego.ContextWithBody[LoginRequest]) (LoginResponse, error) {
	body, err := c.Body()
	if err != nil {
		return LoginResponse{}, err
	}

	response, err := h.loginUseCase.Login(body.toLoginInput())
	if err != nil {
		return LoginResponse{}, err
	}

	return LoginResponse{
		Token: response.Token,
	}, nil
}
