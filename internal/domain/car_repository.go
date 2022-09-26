package domain

import (
	"context"
	"errors"
)

var (
	ErrCarNotFound = errors.New("car not found")
)

type CarRepository interface {
	FindByID(ctx context.Context, id int64) (*Car, error)
	Create(ctx context.Context, c *Car) error
}
