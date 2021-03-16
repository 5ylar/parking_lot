package repository

import (
	"errors"
	"parking_lot/parking"
	"parking_lot/parking/state"
	"strings"
)

type parkingRepository struct{}

func NewParkingRepository() parking.Repository {
	return &parkingRepository{}
}

func (r *parkingRepository) CreateParkLot(num uint16) error {
	if len(state.ParkingLot.Slots) > 0 {
		return errors.New("Can't create parking lots")
	}

	/* create empty slots */
	state.ParkingLot.EmptySlots = num
	state.ParkingLot.Slots = make([]parking.Slot, num)
	for slot := uint16(1); slot <= num; slot++ {
		state.ParkingLot.Slots[slot-1] = parking.Slot{
			SlotNo:  slot,
			IsEmpty: true,
		}
	}

	return nil
}
func (r *parkingRepository) Park(regisNo string, colour string) (uint16, error) {

	if state.ParkingLot.EmptySlots == 0 {
		return 0, errors.New("Sorry, parking lot is full")
	}

	emptyParkingLotIndex, found := getEmptyParkingLotsIndex(state.ParkingLot.Slots)
	if !found {
		return 0, errors.New("Not found")
	}
	slot := &state.ParkingLot.Slots[emptyParkingLotIndex]
	slot.Car = &parking.Car{
		RegisNo: regisNo,
		Colour:  colour,
	}
	slot.IsEmpty = false
	state.ParkingLot.EmptySlots -= 1

	return slot.SlotNo, nil
}
func (r *parkingRepository) Leave(slotNo uint16) error {
	if slotNo == 0 || slotNo > uint16(len(state.ParkingLot.Slots)) {
		return errors.New("Lot's value is invalid")
	}

	slot := &state.ParkingLot.Slots[slotNo-1]
	slot.Car = nil
	slot.IsEmpty = true
	state.ParkingLot.EmptySlots += 1

	return nil
}

func (r *parkingRepository) List(regisNo string, colour string) ([]parking.Slot, error) {
	regisNo = strings.Trim(regisNo, " ")
	colour = strings.Trim(colour, " ")

	if len(state.ParkingLot.Slots) == 0 {
		return nil, nil
	}

	/* search */
	if regisNo != "" || colour != "" {
		var matchedSlots []parking.Slot
		for _, slot := range state.ParkingLot.Slots {
			if slot.IsEmpty || slot.Car == nil {
				continue
			}

			car := *slot.Car
			isMatched := false

			if regisNo != "" && colour != "" {
				if car.RegisNo == regisNo && car.Colour == colour {
					isMatched = true
				}
			} else {
				if regisNo != "" {
					if car.RegisNo == regisNo {
						isMatched = true
					}
				} else if colour != "" {
					if car.Colour == colour {
						isMatched = true
					}
				}
			}

			if isMatched {
				matchedSlots = append(matchedSlots, slot)
			}
		}

		if len(matchedSlots) == 0 {
			return nil, errors.New("Not found")
		}

		return matchedSlots, nil
	}
	/* end search */

	return state.ParkingLot.Slots, nil
}

func getEmptyParkingLotsIndex(slots []parking.Slot) (index int16, found bool) {
	for i := int16(0); i < int16(len(slots)); i++ {
		if slots[i].IsEmpty {
			return i, true
		}
	}

	return 0, false
}
