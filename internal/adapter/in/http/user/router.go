package user

import (
	"net/http"

	"github.com/go-fuego/fuego"
	"github.com/go-fuego/fuego/option"
)

func SetupUserRoutes(s *fuego.Server, userHandler *UserAdapter, authMiddleware func(http.Handler) http.Handler) {
	userGroup := fuego.Group(s, "/users")

	fuego.Post(userGroup, "/register", userHandler.Register)
	fuego.Get(userGroup, "/profile", userHandler.Profile, option.Middleware(authMiddleware))
	fuego.Get(userGroup, "/{id}", userHandler.GetUser, option.Middleware(authMiddleware))
}
