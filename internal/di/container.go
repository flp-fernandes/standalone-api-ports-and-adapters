package di

import (
	"github.com/flp-fernandes/standalone-api-ports-and-adapters/internal/app"
	"github.com/flp-fernandes/standalone-api-ports-and-adapters/internal/infra/postgres"
	"github.com/flp-fernandes/standalone-api-ports-and-adapters/internal/orchestrator"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Container struct {
	Orchestrator *orchestrator.Orchestrator
}

func NewContainer(db *pgxpool.Pool) *Container {
	userRepo := postgres.NewUserRepository(db)
	orderRepo := postgres.NewOrderRepository(db)

	userService := app.NewUserService(userRepo)
	orderService := app.NewOrderService(orderRepo)

	orc := orchestrator.NewOrchestrator(userService, orderService)

	return &Container{
		Orchestrator: orc,
	}
}
