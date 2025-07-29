package domain

import (
	"context"

	"github.com/flp-fernandes/standalone-api-ports-and-adapters/internal/domain/model"
)

type UserRepository interface {
	ListUsers(ctx context.Context) ([]model.User, error)
}
