package main

import (
	"hog-bridge/internal/config"
	"hog-bridge/internal/di"
	"hog-bridge/internal/router"
	"log"
)

func main() {
	cfg := config.LoadConfig()

	h, err := di.Initialize(cfg)
	if err != nil {
		log.Fatal(err)
	}

	r := router.SetupRouter(h)

	err = r.Run(cfg.ServerPort)
	if err != nil {
		log.Fatal(err)
	}
}
