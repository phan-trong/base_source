package usecase

import (
	"context"
	"solid/internal/domain"
)

type BuyCarUseCase interface {
	Handle(ctx context.Context, req *BuyCarRequest) (*BuyCarResponse, error)
}

type BuyCarRequest struct {
	CarID   int64
	BuyerID int64
}
type BuyCarResponse struct {
	Car *domain.Car
}

var _ BuyCarUseCase = &buycarUseCase{}

type buycarUseCase struct {
	carRepo domain.CarRepository
}

// Handle implements BuyCarUseCase
func (uc *buycarUseCase) Handle(ctx context.Context, req *BuyCarRequest) (*BuyCarResponse, error) {
	car, err := uc.carRepo.FindByID(ctx, req.CarID)
	if err != nil {
		return nil, err
	}
	// TODO: Add car to user
	return &BuyCarResponse{car}, nil
}

func NewBuyCarUsecase(carRepo domain.CarRepository) *buycarUseCase {
	return &buycarUseCase{
		carRepo: carRepo,
	}
}
