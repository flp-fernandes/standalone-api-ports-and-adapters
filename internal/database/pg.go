package database

import (
	"context"
	"log"

	"github.com/flp-fernandes/standalone-api-ports-and-adapters/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPGXConnection(cfg config.Config) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(context.Background(), cfg.PostgresUrl)
	if err != nil {
		log.Printf("Failed to create connection pool: %v", err)

		return nil, err
	}

	return pool, nil
}
