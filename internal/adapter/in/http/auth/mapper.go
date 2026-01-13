package auth

import portAuth "github.com/AnataAria/goway/internal/port/in/auth"

func (m *LoginRequest) toLoginInput() *portAuth.LoginRequest {
	return &portAuth.LoginRequest{
		Email:    m.Email,
		Password: m.Password,
	}
}
