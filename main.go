package main

import (
	parkingController "parking_lot/parking/controller"
	parkingRepository "parking_lot/parking/repository"
	parkingUsecase "parking_lot/parking/usecase"
)

func main() {
	parkingRepo := parkingRepository.NewParkingRepository()
	parkingUsecase := parkingUsecase.NewParkingUsecase(parkingRepo)
	parkingController.InitParkingController(parkingUsecase)
}
