package main

import (
	"log"

	"github.com/End-rey/VTBMoreTech5Backend/config"
	"github.com/End-rey/VTBMoreTech5Backend/internal/app"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
