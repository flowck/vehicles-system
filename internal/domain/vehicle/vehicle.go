package vehicle

import (
	"context"
	"errors"
)

type Vehicle struct {
	id     string
	name   string
	models []string
}

func (v *Vehicle) setName(name string) {
	v.name = name
}

func (v *Vehicle) setModels(modes []string) {
	v.models = modes
}

func (v *Vehicle) Name() string {
	return v.name
}

func (v *Vehicle) Models() []string {
	return v.models
}

func (v *Vehicle) Id() string {
	return v.id
}

func NewVehicle(id, name string, models []string) (Vehicle, error) {
	if id == "" {
		return Vehicle{}, errors.New("id cannot be empty")
	}

	if name == "" {
		return Vehicle{}, errors.New("name cannot be empty")
	}

	return Vehicle{
		id:     id,
		name:   name,
		models: models,
	}, nil
}

type Repository interface {
	Find(context.Context, string) (*Vehicle, error)
	FindAll(ctx context.Context) ([]Vehicle, error)
}
