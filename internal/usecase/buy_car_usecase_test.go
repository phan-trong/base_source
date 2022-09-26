package usecase

import (
	"context"
	"solid/internal/domain"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestBuyCarUcHandle(t *testing.T) {

	mctl := gomock.NewController(t)
	carRepo := domain.NewMockCarRepository(mctl)
	uc := NewBuyCarUsecase(carRepo)

	carRepo.EXPECT().FindByID(gomock.Any(), int64(1)).Return(&domain.Car{
		ID:        int64(1),
		ModelName: "",
	}, nil)

	resp, err := uc.Handle(context.TODO(), &BuyCarRequest{
		CarID:   int64(1),
		BuyerID: 123,
	})

	if err != nil {
		t.Fatal()
	}
	if resp.Car.ID != int64(1) {
		t.Fatal()
	}
	t.Log("Success")
}
