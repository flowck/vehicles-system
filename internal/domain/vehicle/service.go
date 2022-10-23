package vehicle

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository}
}

func (s *Service) GetVehicles() ([]Vehicle, error) {
	return nil, nil
}
