package vehicle

import "context"

type Vehicle struct {
	Id     string
	Name   string
	Models []string
}

type Repository interface {
	Find(context.Context, string) (*Vehicle, error)
	FindAll(ctx context.Context) ([]Vehicle, error)
}
