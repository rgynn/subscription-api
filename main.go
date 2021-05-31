package main

import (
	"log"

	"github.com/rgynn/subscription-api/pkg/api"
	"github.com/rgynn/subscription-api/pkg/config"
)

func main() {

	cfg, err := config.NewFromEnv()
	if err != nil {
		log.Fatal(err)
	}

	srv, err := api.NewServerFromConfig(cfg)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Listening on: %s\n", cfg.Port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
