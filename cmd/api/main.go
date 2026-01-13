package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/AnataAria/goway/internal/config"
	"github.com/AnataAria/goway/internal/startup/http"
	"github.com/AnataAria/goway/pkg/postgres"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)

	defer stop()

	pg, err := postgres.New(ctx, postgres.PostgresConfig{
		DSN: fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s?sslmode=%s",
			cfg.Database.User,
			cfg.Database.Password,
			cfg.Database.Host,
			cfg.Database.Port,
			cfg.Database.DBName,
			cfg.Database.SSLMode,
		),
		MaxOpenConns:    cfg.Database.MaxOpenConns,
		MaxIdleConns:    cfg.Database.MaxIdleConns,
		ConnMaxLifetime: time.Minute * 30,
	})

	if err != nil {
		log.Fatalf("failed to create postgres client: %v", err)
	}

	server := http.NewServer(ctx, pg, cfg)
	server.SetupHttp()
	server.Start()
}
