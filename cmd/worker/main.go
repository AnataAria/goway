package main

import (
	"log"

	"github.com/AnataAria/goway/internal/config"
	"github.com/AnataAria/goway/internal/startup/worker"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	if err := worker.StartupWorker(cfg); err != nil {
		log.Fatalf("Failed to start worker: %v", err)
	}
}
