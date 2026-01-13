package http

import (
	"github.com/AnataAria/goway/internal/adapter/in/http/auth"
	"github.com/AnataAria/goway/internal/adapter/in/http/middleware"
	httpUser "github.com/AnataAria/goway/internal/adapter/in/http/user"
	"github.com/AnataAria/goway/internal/adapter/out/persistence/postgres/account"
	postgresUser "github.com/AnataAria/goway/internal/adapter/out/persistence/postgres/user"
	"github.com/AnataAria/goway/internal/adapter/out/security"
	appAuth "github.com/AnataAria/goway/internal/application/auth"
	appUser "github.com/AnataAria/goway/internal/application/user"
)

func (s *Server) SetupHttp() {
	userRepo := postgresUser.NewUserRepository(s.pg.DB)
	accountRepo := account.NewAccountRepository(s.pg.DB)

	passwordHasher := security.NewBcryptPasswordHasher()
	tokenGenerator := security.NewJWTTokenGenerator(s.cfg.JWT.SecretKey)
	authMiddleware := middleware.JWTAuth(tokenGenerator)

	userUseCase := appUser.New(userRepo)
	userHandler := httpUser.NewUserAdapter(userUseCase)

	loginUseCase := appAuth.New(accountRepo, passwordHasher, tokenGenerator)
	authHandler := auth.NewAuthAdapter(loginUseCase)

	httpUser.SetupUserRoutes(s.fuego, userHandler, authMiddleware)
	auth.SetupAuthRoutes(s.fuego, authHandler)

}
