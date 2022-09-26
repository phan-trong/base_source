package main

import (
	"net/http"
	"solid/internal/adapter"
	httpport "solid/internal/ports/http"
	"solid/internal/usecase"
)

func main() {
	carRepo := adapter.NewInMemCarRepo()
	buycarUseCase := usecase.NewBuyCarUsecase(carRepo)

	uc := usecase.Usecase{
		BuyCar: buycarUseCase,
	}

	httpHandler := httpport.NewBuyCarHandler(uc, carRepo)

	mux := http.NewServeMux()
	mux.HandleFunc("/buy-car", httpHandler.BuyCar)
	mux.HandleFunc("/create-car", httpHandler.CreateCar)
	http.ListenAndServe(":8088", mux)
}
