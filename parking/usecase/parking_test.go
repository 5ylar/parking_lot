package usecase_test

import (
	"fmt"
	parking "parking_lot/parking"
	parkingRepository "parking_lot/parking/repository"
	parkingUsecase "parking_lot/parking/usecase"
	"testing"
)

var usecaseInstance parking.Usecase

func getUsecaseInstance() parking.Usecase {
	if usecaseInstance != nil {
		return usecaseInstance
	}

	usecaseInstance = parkingUsecase.NewParkingUsecase(parkingRepository.NewParkingRepository())
	return usecaseInstance
}

func TestCreateParkingLots(t *testing.T) {
	usecase := getUsecaseInstance()
	const parkingLotsNum = 6
	msg, err := usecase.CreateParkLot(parkingLotsNum)
	if err != nil {
		t.Fatal(err)
	}
	if msg != fmt.Sprintf("Created a parking lot with %d slots", parkingLotsNum) {
		t.Fatal("Create parking lots not successfully")
	}
}

func TestPark(t *testing.T) {
	usecase := getUsecaseInstance()
	const regisNo = "KA-01-BB-0001"
	const color = "Black"
	msg, err := usecase.Park(regisNo, color)
	if err != nil {
		t.Fatal(err)
	}
	if msg != "Allocated slot number: 1" {
		t.Fatal("Park not successfully")
	}
}

func TestPark2(t *testing.T) {
	usecase := getUsecaseInstance()
	const regisNo = "KA-01-BB-0002"
	const color = "Red"
	msg, err := usecase.Park(regisNo, color)
	if err != nil {
		t.Fatal(err)
	}
	if msg != "Allocated slot number: 2" {
		t.Fatal("Park not successfully")
	}
}

func TestLeave(t *testing.T) {
	usecase := getUsecaseInstance()
	const slot = 2
	msg, err := usecase.Leave(slot)
	if err != nil {
		t.Fatal(err)
	}
	if msg != fmt.Sprintf("Slot number %d is free", slot) {
		t.Fatal("Leave not successfully")
	}
}

func TestStatus(t *testing.T) {
	usecase := getUsecaseInstance()
	msg, err := usecase.Status()
	if err != nil {
		t.Fatal(err)
	}
	expectedMsg := "Slot No.    Registration No    Colour\n1           KA-01-BB-0001      Black"
	if msg != expectedMsg {
		t.Fatal("Wrong status")
	}
}

func TestFindRegisNoForCarByColour(t *testing.T) {
	usecase := getUsecaseInstance()
	const color = "Black"
	msg, err := usecase.FindRegisNoForCarByColour(color)
	if err != nil {
		t.Fatal(err)
	}
	if msg != "KA-01-BB-0001" {
		t.Fatal("Not found")
	}
}

func TestFindSlotNumberByColour(t *testing.T) {
	usecase := getUsecaseInstance()
	const color = "Black"
	msg, err := usecase.FindSlotNumberByColour(color)
	if err != nil {
		t.Fatal(err)
	}
	if msg != "1" {
		t.Fatal("Not found")
	}
}

func TestFindSlotNumberByRegisNo(t *testing.T) {
	usecase := getUsecaseInstance()
	const regisNo = "KA-01-BB-0001"
	msg, err := usecase.FindSlotNumberByRegisNo(regisNo)
	if err != nil {
		t.Fatal(err)
	}
	if msg != "1" {
		t.Fatal("Not found")
	}
}

//

func TestFindRegisNoForCarByColourNotFound(t *testing.T) {
	usecase := getUsecaseInstance()
	const color = "Green"
	_, err := usecase.FindRegisNoForCarByColour(color)
	if err == nil || err.Error() != "Not found" {
		t.Fatal("Should be \"not found\"")
	}
}

func TestFindSlotNumberByColourNotFound(t *testing.T) {
	usecase := getUsecaseInstance()
	const color = "Green"
	_, err := usecase.FindSlotNumberByColour(color)
	if err == nil || err.Error() != "Not found" {
		t.Fatal("Should be \"not found\"")
	}
}

func TestFindSlotNumberByRegisNoNotFound(t *testing.T) {
	usecase := getUsecaseInstance()
	const regisNo = "KA-01-BB-0003"
	_, err := usecase.FindSlotNumberByRegisNo(regisNo)
	if err == nil || err.Error() != "Not found" {
		t.Fatal("Should be \"not found\"")
	}
}

func TestParkButFull(t *testing.T) {
	usecase := getUsecaseInstance()
	const regisNo = "KA-01-BB-0002"
	const color = "Blue"
	for i := 0; i < 5; i++ {
		usecase.Park(regisNo, color)
	}
	_, err := usecase.Park(regisNo, color)
	if err == nil || err.Error() != "Sorry, parking lot is full" {
		t.Fatal("Should be \"Sorry, parking lot is full\"")
	}
}
