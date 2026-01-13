package user

import (
	"github.com/go-fuego/fuego"
)

func SetupUserRoutes(s *fuego.Server, userHandler *UserHandler) {
	userGroup := fuego.Group(s, "/users")

	fuego.Post(userGroup, "/register", userHandler.Register)
	fuego.Get(userGroup, "/profile", userHandler.Profile)
	fuego.Get(userGroup, "/{id}", userHandler.GetUser)

}
