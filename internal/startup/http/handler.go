package http

import (
	"github.com/AnataAria/goway/internal/adapter/in/http/auth"
	httpUser "github.com/AnataAria/goway/internal/adapter/in/http/user"
	"github.com/AnataAria/goway/internal/adapter/out/persistence/postgres/account"
	postgresUser "github.com/AnataAria/goway/internal/adapter/out/persistence/postgres/user"
	"github.com/AnataAria/goway/internal/adapter/out/security"
)

func (s *Server) SetupHttp() {
	userRepo := postgresUser.NewUserRepository(s.pg.DB)
	accountRepo := account.NewAccountRepository(s.pg.DB)

	passwordHasher := security.NewBcryptPasswordHasher()
	tokenGenerator := security.NewJWTTokenGenerator(s.cfg.JWT.SecretKey)

	userHandler := httpUser.NewUserAdapter(userRepo)
	authHandler := auth.NewAuthAdapter(accountRepo, passwordHasher, tokenGenerator)

	httpUser.SetupUserRoutes(s.fuego, userHandler)
	auth.SetupAuthRoutes(s.fuego, authHandler)

}
