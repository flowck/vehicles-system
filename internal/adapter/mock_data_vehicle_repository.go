package adapter

import (
	"context"
	"vehicles-system/internal/domain/vehicle"
)

type MockDataVehicleRepository struct {
}

func NewMockDataVehicleRepository() MockDataVehicleRepository {
	return MockDataVehicleRepository{}
}

func (r MockDataVehicleRepository) Find(ctx context.Context, id string) (*vehicle.Vehicle, error) {
	return nil, nil
}

func (r MockDataVehicleRepository) FindAll(ctx context.Context) ([]vehicle.Vehicle, error) {
	return nil, nil
}
