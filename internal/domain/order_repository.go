package domain

import (
	"context"

	"github.com/flp-fernandes/standalone-api-ports-and-adapters/internal/domain/model"
)

type OrderRepository interface {
	ListOrders(ctx context.Context) ([]model.Order, error)
}
