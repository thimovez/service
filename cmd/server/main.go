package main

import (
	"github.com/thimovez/service/config"
	"github.com/thimovez/service/internal/app"
	"log"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)
}
