package http

import (
	"context"

	"github.com/AnataAria/goway/internal/config"
	"github.com/AnataAria/goway/pkg/postgres"
	"github.com/go-fuego/fuego"
)

type Server struct {
	ctx   context.Context
	pg    *postgres.DB
	cfg   *config.Config
	fuego *fuego.Server
}

func NewServer(ctx context.Context, pg *postgres.DB, config *config.Config) *Server {
	s := fuego.NewServer()
	return &Server{
		ctx:   ctx,
		pg:    pg,
		cfg:   config,
		fuego: s,
	}
}
