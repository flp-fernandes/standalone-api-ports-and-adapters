package app

import (
	"context"
	"log"
	"time"

	"github.com/flp-fernandes/standalone-api-ports-and-adapters/internal/domain"
)

type OrderService struct {
	repo domain.OrderRepository
}

func NewOrderService(repo domain.OrderRepository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) PrintAllOrders(ctx context.Context) error {
	log.Println("Started OrderService.PrintAllOrders")

	orders, err := s.repo.ListOrders(ctx)
	if err != nil {
		return err
	}

	for _, o := range orders {
		log.Printf("Order: %d (User: %d) - R$%.2f", o.ID, o.UserID, o.Amount)
	}

	time.Sleep(15 * time.Second)

	return nil
}
