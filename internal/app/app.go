package app

import (
	"fmt"
	"vehicles-system/internal/adapters"
	"vehicles-system/internal/domain/vehicle"
)

type App struct {
	VehicleService *vehicle.Service
}

func NewApp() *App {
	vehicleRepository, err := adapters.NewMockDataVehicleRepository()
	if err != nil {
		panic(fmt.Sprintf("cannot create repository vehicleRepository: %v", err))
	}

	vehicleService := vehicle.NewService(vehicleRepository)

	return &App{VehicleService: vehicleService}
}
