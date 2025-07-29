package database

import (
	"context"

	"github.com/flp-fernandes/standalone-api-ports-and-adapters/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPGXConnection(cfg config.Config) (*pgxpool.Pool, error) {
	return pgxpool.New(context.Background(), cfg.PostgresUrl)
}
