package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/AnataAria/goway/internal/config"
	"github.com/AnataAria/goway/internal/startup/http"
	"github.com/AnataAria/goway/pkg/postgres"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	ctx := context.Background()

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.DBName,
		cfg.Database.SSLMode,
	)

	pgConfig := postgres.PostgresConfig{
		DSN:             dsn,
		MaxOpenConns:    cfg.Database.MaxOpenConns,
		MaxIdleConns:    cfg.Database.MaxIdleConns,
		ConnMaxLifetime: 5 * time.Minute,
	}

	pg, err := postgres.New(ctx, pgConfig)
	if err != nil {
		log.Fatalf("Postgres error: %s", err)
	}
	defer pg.Close()

	server := http.NewServer(ctx, pg, cfg)
	server.SetupHttp()
	server.Start()
}
