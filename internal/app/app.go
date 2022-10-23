package app

import (
	"vehicles-system/internal/adapter"
	"vehicles-system/internal/domain/vehicle"
)

type App struct {
	VehicleService *vehicle.Service
}

func NewApp() *App {
	vehicleRepository := adapter.NewMockDataVehicleRepository()
	vehicleService := vehicle.NewService(vehicleRepository)

	return &App{VehicleService: vehicleService}
}
