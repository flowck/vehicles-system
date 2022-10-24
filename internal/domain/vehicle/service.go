package vehicle

import (
	"context"
	"strings"
)

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository}
}

func (s *Service) GetVehicles(ctx context.Context) ([]Vehicle, error) {
	return s.repository.FindAll(ctx)
}

func (s *Service) GetVehicle(ctx context.Context, id string) (*Vehicle, error) {
	return s.repository.Find(ctx, strings.ToLower(id))
}
