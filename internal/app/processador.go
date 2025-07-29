package app

import (
	"context"
	"log"

	"github.com/flp-fernandes/standalone-api-ports-and-adapters/internal/domain"
)

type Service struct {
	repo domain.UserRepository
}

func NewService(repo domain.UserRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Process(ctx context.Context) error {
	users, err := s.repo.ListUsers(ctx)
	if err != nil {
		return err
	}

	for _, user := range users {
		log.Printf("Usuário: %d - %s", user.ID, user.Name)
	}

	return nil
}
