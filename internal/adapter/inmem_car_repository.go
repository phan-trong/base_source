package adapter

import (
	"context"
	"fmt"
	"math/rand"
	"solid/internal/domain"
)

var _ domain.CarRepository = &carRepositoryInMem{}

type carRepositoryInMem struct {
	m map[int64]*domain.Car
}

func NewInMemCarRepo() *carRepositoryInMem {
	return &carRepositoryInMem{
		m: make(map[int64]*domain.Car),
	}
}

// Create implements domain.CarRepository
func (r *carRepositoryInMem) Create(ctx context.Context, c *domain.Car) error {
	id := rand.Int31()
	c.ID = int64(id)
	r.m[c.ID] = c
	return nil
}

// FindByID implements domain.CarRepository
func (r *carRepositoryInMem) FindByID(ctx context.Context, id int64) (*domain.Car, error) {
	fmt.Printf("%v\n", r.m)
	if car, ok := r.m[id]; ok {
		return car, nil
	}
	return nil, domain.ErrCarNotFound
}
