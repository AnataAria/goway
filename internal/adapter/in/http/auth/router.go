package auth

import (
	"github.com/go-fuego/fuego"
)

func SetupAuthRoutes(s *fuego.Server, authHandler *AuthAdapter) {
	authGroup := fuego.Group(s, "/auth")

	fuego.Post(authGroup, "/login", authHandler.Login)
}
