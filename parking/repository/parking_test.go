package repository_test

import (
	parkingRepository "parking_lot/parking/repository"
	parkingState "parking_lot/parking/state"
	"testing"
)

func TestCreateParkLots(t *testing.T) {
	_parkingRepository := parkingRepository.NewParkingRepository()
	const parkSlotsNum = 6
	if err := _parkingRepository.CreateParkLot(parkSlotsNum); err != nil {
		t.Fatal(err)
	}

	if len(parkingState.ParkingLot.Slots) != parkSlotsNum {
		t.Fatal("Create park lots not successfully")
	}
}

func TestPark(t *testing.T) {
	_parkingRepository := parkingRepository.NewParkingRepository()
	const regisNo = "KA-01-BB-0001"
	const color = "Black"
	if _, err := _parkingRepository.Park(regisNo, color); err != nil {
		t.Fatal(err)
	}
	slot := parkingState.ParkingLot.Slots[0]

	if slot.IsEmpty ||
		slot.Car.RegisNo != regisNo ||
		slot.Car.Colour != color {
		t.Fatal("Park not successfully")
	}
}

func TestPark2(t *testing.T) {
	_parkingRepository := parkingRepository.NewParkingRepository()
	const regisNo = "KA-01-BB-0002"
	const color = "Red"
	if _, err := _parkingRepository.Park(regisNo, color); err != nil {
		t.Fatal(err)
	}
	slot := parkingState.ParkingLot.Slots[1]

	if slot.IsEmpty ||
		slot.Car.RegisNo != regisNo ||
		slot.Car.Colour != color {
		t.Fatal("Park not successfully")
	}
}

func TestLeave(t *testing.T) {
	_parkingRepository := parkingRepository.NewParkingRepository()
	const slotNo = 2
	if err := _parkingRepository.Leave(slotNo); err != nil {
		t.Fatal(err)
	}
	slot := parkingState.ParkingLot.Slots[1]

	if !slot.IsEmpty {
		t.Fatal("Leave not successfully")
	}
}

func TestList(t *testing.T) {
	_parkingRepository := parkingRepository.NewParkingRepository()
	slots, err := _parkingRepository.List("", "")
	if err != nil {
		t.Fatal(err)
	}
	stateParkLotSlots := parkingState.ParkingLot.Slots

	if len(slots) != len(stateParkLotSlots) {
		t.Fatal("Wrong information ")
	}

	for slotIndex, slot := range slots {
		if slotIndex == 0 {
			if slot.IsEmpty {
				t.Fatal("Wrong information ")
			}
		} else {
			if !slot.IsEmpty {
				t.Fatal("Wrong information ")
			}
		}
	}
}

func TestFindByColour(t *testing.T) {
	_parkingRepository := parkingRepository.NewParkingRepository()
	const color = "Black"
	slots, err := _parkingRepository.List("", color)
	if err != nil {
		t.Fatal(err)
	}

	if len(slots) != 1 {
		t.Fatal("Wrong information ")
	}

	if slots[0].IsEmpty ||
		slots[0].Car.Colour != color {
		t.Fatal("Wrong information ")
	}
}

func TestFindByRegisNo(t *testing.T) {
	_parkingRepository := parkingRepository.NewParkingRepository()
	const regisNo = "KA-01-BB-0001"
	slots, err := _parkingRepository.List(regisNo, "")
	if err != nil {
		t.Fatal(err)
	}

	if len(slots) != 1 {
		t.Fatal("Wrong information ")
	}

	if slots[0].IsEmpty ||
		slots[0].Car.RegisNo != regisNo {
		t.Fatal("Wrong information ")
	}
}

func TestFindByRegisNoAndColor(t *testing.T) {
	_parkingRepository := parkingRepository.NewParkingRepository()
	const regisNo = "KA-01-BB-0001"
	const color = "Black"
	slots, err := _parkingRepository.List(regisNo, color)
	if err != nil {
		t.Fatal(err)
	}

	if len(slots) != 1 {
		t.Fatal("Wrong information ")
	}

	if slots[0].IsEmpty ||
		slots[0].Car.Colour != color ||
		slots[0].Car.RegisNo != regisNo {
		t.Fatal("Wrong information ")
	}
}

func TestFindByColourNotFound(t *testing.T) {
	_parkingRepository := parkingRepository.NewParkingRepository()
	const color = "Red"
	_, err := _parkingRepository.List("", color)
	if err == nil || err.Error() != "Not found" {
		t.Fatal("Should be \"not found\"")
	}
}

func TestFindByRegisNoNotFound(t *testing.T) {
	_parkingRepository := parkingRepository.NewParkingRepository()
	const regisNo = "KA-01-BB-0002"
	_, err := _parkingRepository.List(regisNo, "")
	if err == nil || err.Error() != "Not found" {
		t.Fatal("Should be \"not found\"")
	}
}

func TestFindByRegisNoAndColorNotFound(t *testing.T) {
	_parkingRepository := parkingRepository.NewParkingRepository()
	const regisNo = "KA-01-BB-0002"
	const color = "Black"
	_, err := _parkingRepository.List(regisNo, color)
	if err == nil || err.Error() != "Not found" {
		t.Fatal("Should be \"not found\"")
	}
}

func TestParkButFull(t *testing.T) {
	_parkingRepository := parkingRepository.NewParkingRepository()
	const regisNo = "KA-01-BB-0002"
	const color = "Red"

	for i := 0; i < 5; i++ {
		_parkingRepository.Park(regisNo, color)
	}

	_, err := _parkingRepository.Park(regisNo, color)
	if err == nil || err.Error() != "Sorry, parking lot is full" {
		t.Fatal("Should be \"Sorry, parking lot is full\"")
	}
}
