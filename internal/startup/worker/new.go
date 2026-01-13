package worker

import (
	"log"

	"github.com/AnataAria/goway/internal/config"
)

func StartupWorker(cfg *config.Config) error {
	log.Printf("Starting worker...")
	log.Printf("Worker started successfully")
	return nil
}
