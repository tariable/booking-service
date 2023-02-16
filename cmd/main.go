package main

import (
	"bookingService/config"
	"bookingService/internal/app"
	"log"
)

func main() {
	// Load config
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run application
	app.Run(cfg)
}
