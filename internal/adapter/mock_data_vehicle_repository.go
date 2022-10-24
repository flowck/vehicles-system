/**
This is supposed to be a mock repository that implements the interface vehicle.Repository, therefore is can be discarded
later. ;)
*/

package adapter

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"vehicles-system/internal/domain/vehicle"
)

type mockDataVehicleRepository struct {
	data []vehicle.Vehicle
}

type mockVehicle struct {
	Brand  string   `json:"brand"`
	Models []string `json:"models"`
}

func NewMockDataVehicleRepository() (mockDataVehicleRepository, error) {
	workdir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	
	file, err := os.ReadFile(fmt.Sprintf("%s/internal/adapter/cars.json", workdir))

	if err != nil {
		panic(err)
	}

	var mockData []mockVehicle
	if err = json.Unmarshal(file, &mockData); err != nil {
		panic(err)
	}

	data := make([]vehicle.Vehicle, len(mockData))
	// Map mock data
	for idx, item := range mockData {
		v, err := vehicle.NewVehicle(fmt.Sprintf("%s-%d", strings.ToLower(item.Brand), idx), item.Brand, item.Models)
		if err != nil {
			return mockDataVehicleRepository{}, err
		}

		data[idx] = v
	}

	return mockDataVehicleRepository{data}, nil
}

func (r mockDataVehicleRepository) Find(ctx context.Context, id string) (*vehicle.Vehicle, error) {
	return nil, nil
}

func (r mockDataVehicleRepository) FindAll(ctx context.Context) ([]vehicle.Vehicle, error) {
	return r.data, nil
}
