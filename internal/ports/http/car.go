package http

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"solid/internal/domain"
	"solid/internal/usecase"
	"strconv"
)

type carHandler struct {
	uc      usecase.Usecase
	carRepo domain.CarRepository
}

func NewBuyCarHandler(uc usecase.Usecase, carRepo domain.CarRepository) *carHandler {
	return &carHandler{
		uc:      uc,
		carRepo: carRepo,
	}
}

type ErrorResponse struct {
	Code    int
	Message string
}

type CarDTO struct {
	ID        int64
	ModelName string
	IsInUse   bool
}

// ServeHTTP implements http.Handler
func (h *carHandler) BuyCar(w http.ResponseWriter, req *http.Request) {
	var (
		carID  = req.Form.Get("car_id")
		userID = req.Form.Get("user_id")
	)
	carIDi64, _ := strconv.ParseInt(carID, 10, 64)
	userIDi64, _ := strconv.ParseInt(userID, 10, 64)

	resp, err := h.uc.BuyCar.Handle(req.Context(), &usecase.BuyCarRequest{
		CarID:   carIDi64,
		BuyerID: userIDi64,
	})
	if err != nil {
		// TODO handle error
		json.NewEncoder(w).Encode(&ErrorResponse{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	// Mapper
	carDTO := &CarDTO{
		ID:        resp.Car.ID,
		ModelName: resp.Car.ModelName,
		IsInUse:   resp.Car.IsInUse(),
	}

	json.NewEncoder(w).Encode(carDTO)
}

func (h *carHandler) CreateCar(w http.ResponseWriter, r *http.Request) {
	car := &domain.Car{
		ModelName: fmt.Sprintf("Car[%d]", rand.Int31()),
	}
	_ = h.carRepo.Create(r.Context(), car)
	json.NewEncoder(w).Encode(car)
}
