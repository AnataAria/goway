package auth

import (
	"github.com/go-fuego/fuego"
)

func SetupAuthRoutes(s *fuego.Server, authHandler *AuthHandler) {
	authGroup := fuego.Group(s, "/auth")

	fuego.Post(authGroup, "/login", authHandler.Login)
}
