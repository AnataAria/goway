package security

import (
	security "github.com/AnataAria/goway/internal/port/out/security"
	"golang.org/x/crypto/bcrypt"
)

type BcryptPasswordHasher struct{}

func NewBcryptPasswordHasher() security.PasswordHasher {
	return &BcryptPasswordHasher{}
}

func (h *BcryptPasswordHasher) Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (h *BcryptPasswordHasher) Verify(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
