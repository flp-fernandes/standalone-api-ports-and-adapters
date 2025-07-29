package app

import (
	"context"
	"log"
	"time"

	"github.com/flp-fernandes/standalone-api-ports-and-adapters/internal/domain"
)

type UserService struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) PrintAllUsers(ctx context.Context) error {
	log.Println("Started UserService.PrintAllOrders")

	users, err := s.repo.ListUsers(ctx)
	if err != nil {
		return err
	}

	for _, u := range users {
		log.Printf("User: %d - %s", u.ID, u.Name)
	}

	time.Sleep(15 * time.Second)

	return nil
}
