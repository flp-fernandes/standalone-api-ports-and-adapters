package postgres

import (
	"context"

	"github.com/flp-fernandes/standalone-api-ports-and-adapters/internal/domain"
	"github.com/flp-fernandes/standalone-api-ports-and-adapters/internal/domain/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type orderRepository struct {
	db *pgxpool.Pool
}

func NewOrderRepository(db *pgxpool.Pool) domain.OrderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) ListOrders(ctx context.Context) ([]model.Order, error) {
	query := `SELECT id, user_id, amount FROM orders`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []model.Order
	for rows.Next() {
		var o model.Order
		if err := rows.Scan(&o.ID, &o.UserID, &o.Amount); err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}

	return orders, nil
}
