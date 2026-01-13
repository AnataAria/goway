package security

import (
	"time"

	security "github.com/AnataAria/goway/internal/port/out/security"
	"github.com/golang-jwt/jwt/v5"
)

type JWTTokenGenerator struct {
	secretKey string
}

func NewJWTTokenGenerator(secretKey string) security.TokenGenerator {
	return &JWTTokenGenerator{secretKey: secretKey}
}

func (g *JWTTokenGenerator) Generate(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"iat":     time.Now().Unix(),
	})

	return token.SignedString([]byte(g.secretKey))
}

func (g *JWTTokenGenerator) Validate(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(g.secretKey), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if userID, ok := claims["user_id"].(string); ok {
			return userID, nil
		}
	}

	return "", jwt.ErrSignatureInvalid
}
