package http

import (
	"context"
	"fmt"

	"github.com/AnataAria/goway/internal/adapter/in/http/middleware"
	"github.com/AnataAria/goway/internal/config"
	"github.com/AnataAria/goway/pkg/postgres"
	"github.com/AnataAria/goway/pkg/validator"
	"github.com/go-fuego/fuego"
)

type Server struct {
	ctx   context.Context
	pg    *postgres.DB
	cfg   *config.Config
	fuego *fuego.Server
}

func NewServer(ctx context.Context, pg *postgres.DB, config *config.Config) *Server {
	limiter := middleware.NewRateLimiter(20, 5)

	s := fuego.NewServer(
		fuego.WithAddr(fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)),
		fuego.WithValidator(validator.New().Validator()),
		fuego.WithEngineOptions(
			fuego.WithErrorHandler(errorHandler),
		),
	)

	fuego.Use(s, middleware.Logger)
	fuego.Use(s, limiter.Middleware)

	return &Server{
		ctx:   ctx,
		pg:    pg,
		cfg:   config,
		fuego: s,
	}
}
