package controllor_test

import (
	"io/ioutil"
	"os"
	parkingController "parking_lot/parking/controller"
	parkingRepository "parking_lot/parking/repository"
	parkingUsecase "parking_lot/parking/usecase"
	"strings"
	"testing"
)

func captureOutput(f func()) string {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	f()
	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	return strings.Trim(string(out), "\n")
}

var parkingControllerInstance *parkingController.ParkingController

func getControllerInstance() *parkingController.ParkingController {
	if parkingControllerInstance != nil {
		return parkingControllerInstance
	}
	parkingRepo := parkingRepository.NewParkingRepository()
	parkingUsecase := parkingUsecase.NewParkingUsecase(parkingRepo)
	parkingControllerInstance = parkingController.NewParkingController(parkingUsecase)
	return parkingControllerInstance
}

func testLogOutputWrapper(t *testing.T, command, expectResult string) error {
	_parkingController := getControllerInstance()
	var err error
	result := captureOutput(func() {
		_, err = _parkingController.ProcessCommand(command)
	})

	if err != nil {
		result = err.Error()
	}

	if result != expectResult {
		t.Fatalf("Expected \"%s\", got \"%s\"", expectResult, result)
	}

	return nil
}

func TestCreateParkLots(t *testing.T) {
	testLogOutputWrapper(
		t,
		"create_parking_lot 6",
		"Created a parking lot with 6 slots",
	)
}

func TestPark(t *testing.T) {
	testLogOutputWrapper(
		t,
		"park KA-01-BB-0001 Black",
		"Allocated slot number: 1",
	)
}

func TestPark2(t *testing.T) {
	testLogOutputWrapper(
		t,
		"park KA-01-BB-0002 Red",
		"Allocated slot number: 2",
	)
}

func TestLeave(t *testing.T) {
	testLogOutputWrapper(
		t,
		"leave 2",
		"Slot number 2 is free",
	)
}

func TestGetStatus(t *testing.T) {
	testLogOutputWrapper(
		t,
		"status",
		"Slot No.    Registration No    Colour\n1           KA-01-BB-0001      Black",
	)
}

func TestFindRestNoByColour(t *testing.T) {
	testLogOutputWrapper(
		t,
		"registration_numbers_for_cars_with_colour Black",
		"KA-01-BB-0001",
	)
}

func TestFindSlotNumberByColour(t *testing.T) {
	testLogOutputWrapper(
		t,
		"slot_numbers_for_cars_with_colour Black",
		"1",
	)
}

func TestFindSlotNumberByRegisNo(t *testing.T) {
	testLogOutputWrapper(
		t,
		"slot_number_for_registration_number KA-01-BB-0001",
		"1",
	)
}

func TestFindRestNoByColourNotFound(t *testing.T) {
	testLogOutputWrapper(
		t,
		"registration_numbers_for_cars_with_colour Green",
		"Not found",
	)
}

func TestFindSlotNumberByColourNotFound(t *testing.T) {
	testLogOutputWrapper(
		t,
		"slot_numbers_for_cars_with_colour Green",
		"Not found",
	)
}

func TestFindSlotNumberByRegisNoNotFound(t *testing.T) {
	testLogOutputWrapper(
		t,
		"slot_number_for_registration_number KA-01-BB-0003",
		"Not found",
	)
}

func TestParkButFull(t *testing.T) {
	_parkingController := getControllerInstance()
	for i := 0; i < 5; i++ {
		_parkingController.ProcessCommand("park KA-01-BB-0002 Blue")
	}

	testLogOutputWrapper(
		t,
		"park KA-01-BB-0002 Green",
		"Sorry, parking lot is full",
	)
}
