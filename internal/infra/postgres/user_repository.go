package postgres

import (
	"context"

	"github.com/flp-fernandes/standalone-api-ports-and-adapters/internal/domain"
	"github.com/flp-fernandes/standalone-api-ports-and-adapters/internal/domain/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type userRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) domain.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) ListUsers(ctx context.Context) ([]model.User, error) {
	query := `SELECT id, name FROM users`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var u model.User
		if err := rows.Scan(&u.ID, &u.Name); err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	return users, nil
}
