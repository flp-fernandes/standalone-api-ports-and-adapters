package orchestrator

import (
	"context"
	"log"

	"github.com/flp-fernandes/standalone-api-ports-and-adapters/internal/app"
)

type Orchestrator struct {
	UserService  *app.UserService
	OrderService *app.OrderService
}

func NewOrchestrator(userService *app.UserService, orderService *app.OrderService) *Orchestrator {
	return &Orchestrator{
		UserService:  userService,
		OrderService: orderService,
	}
}

func (o *Orchestrator) Run(ctx context.Context) error {
	log.Println("Inicializando processamento centralizado...")

	if err := o.UserService.PrintAllUsers(ctx); err != nil {
		return err
	}

	if err := o.OrderService.PrintAllOrders(ctx); err != nil {
		return err
	}

	log.Println("Processamento concluído com sucesso.")

	return nil
}
