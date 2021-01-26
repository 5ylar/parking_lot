package usecase

import (
	"errors"
	"fmt"
	"parking_lot/parking"
	"strconv"
	"strings"
)

type parkingUsecase struct {
	parkingRepo parking.Repository
}

func NewParkingUsecase(_parkingRepo parking.Repository) parking.Usecase {
	return &parkingUsecase{
		parkingRepo: _parkingRepo,
	}
}

func (u *parkingUsecase) CreateParkLot(num uint16) (string, error) {
	if num == 0 {
		return "", errors.New("Required number of parking lots")
	}

	err := u.parkingRepo.CreateParkLot(num)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Created a parking lot with %d slots", num), nil
}
func (u *parkingUsecase) Park(regisNo string, colour string) (string, error) {
	slot, err := u.parkingRepo.Park(regisNo, colour)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Allocated slot number: %d", slot), nil
}
func (u *parkingUsecase) Leave(slot uint16) (string, error) {
	err := u.parkingRepo.Leave(slot)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Slot number %d is free", slot), nil
}
func (u *parkingUsecase) FindRegisNoForCarByColour(colour string) (string, error) {
	slots, err := u.parkingRepo.List("", colour)
	if err != nil {
		return "", err
	}

	var regisNoListStr string
	for _, slot := range slots {
		if slot.Car == nil {
			continue
		}

		car := *slot.Car
		regisNoListStr += car.RegisNo + ", "
	}

	regisNoListStr = strings.TrimSuffix(regisNoListStr, ", ")

	return regisNoListStr, nil
}
func (u *parkingUsecase) FindSlotNumberByColour(colour string) (string, error) {
	slots, err := u.parkingRepo.List("", colour)
	if err != nil {
		return "", err
	}

	var slotListStr string
	for _, slot := range slots {
		slotStr := strconv.FormatInt(int64(slot.SlotNo), 10)
		slotListStr += slotStr + ", "
	}

	slotListStr = strings.TrimSuffix(slotListStr, ", ")

	return slotListStr, nil
}
func (u *parkingUsecase) FindSlotNumberByRegisNo(regisNo string) (string, error) {
	slots, err := u.parkingRepo.List(regisNo, "")
	if err != nil {
		return "", err
	}

	var slotListStr string
	for _, slot := range slots {
		slotStr := strconv.FormatInt(int64(slot.SlotNo), 10)
		slotListStr += slotStr + ", "
	}

	slotListStr = strings.TrimSuffix(slotListStr, ", ")

	return slotListStr, nil
}
func (u *parkingUsecase) Status() (string, error) {
	slots, err := u.parkingRepo.List("", "")
	if err != nil {
		return "", err
	}

	statusStr := "Slot No.    Registration No    Colour"
	for _, slot := range slots {
		var regisNo string
		var colour string

		if slot.IsEmpty || slot.Car == nil {
			continue
		}

		car := *slot.Car
		regisNo = car.RegisNo
		colour = car.Colour

		statusStr += fmt.Sprintf("\n%d           %s      %s", slot.SlotNo, regisNo, colour)
	}

	return statusStr, nil
}
