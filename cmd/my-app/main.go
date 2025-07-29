package main

import (
	"context"
	"log"

	"github.com/flp-fernandes/standalone-api-ports-and-adapters/internal/config"
	"github.com/flp-fernandes/standalone-api-ports-and-adapters/internal/database"
	"github.com/flp-fernandes/standalone-api-ports-and-adapters/internal/di"
)

func main() {
	ctx := context.Background()

	cfg := config.Load()

	db, err := database.NewPGXConnection(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	container := di.NewContainer(db)

	if err := container.Orchestrator.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
