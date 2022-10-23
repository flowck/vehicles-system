package vehicle

import "context"

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository}
}

func (s *Service) GetVehicles(ctx context.Context) ([]Vehicle, error) {
	return s.repository.FindAll(ctx)
}
