package controllor

import (
	"errors"
	"fmt"
	"os"
	"parking_lot/parking"
	"strconv"
	"strings"
)

type ParkingController struct {
	parkingUsecase parking.Usecase
}

func (c *ParkingController) createParkingLot(args []string) error {
	if len(args) == 0 {
		return errors.New("Required number of parking lot")
	}
	nLot, err := strconv.Atoi(args[0])
	if err != nil {
		return errors.New("Wrong argument for parking lot")
	}

	msg, err := c.parkingUsecase.CreateParkLot(uint16(nLot))
	if err != nil {
		return err
	}
	fmt.Println(msg)

	return nil
}

func (c *ParkingController) park(args []string) error {
	if len(args) < 2 {
		return errors.New("Required at least 2 parameters (registeration number, colour)")
	}

	msg, err := c.parkingUsecase.Park(args[0], args[1])
	if err != nil {
		return err
	}
	fmt.Println(msg)

	return nil
}

func (c *ParkingController) leave(args []string) error {
	if len(args) == 0 {
		return errors.New("Required number of parking lot")
	}
	nLot, err := strconv.Atoi(args[0])
	if err != nil {
		return errors.New("Wrong argument for parking lot")
	}

	msg, err := c.parkingUsecase.Leave(uint16(nLot))
	if err != nil {
		return err
	}
	fmt.Println(msg)
	return nil
}

func (c *ParkingController) status(args []string) error {
	msg, err := c.parkingUsecase.Status()
	if err != nil {
		return err
	}
	fmt.Println(msg)
	return nil
}

func (c *ParkingController) findRegisNoForCarByColour(args []string) error {
	if len(args) == 0 {
		return errors.New("Required colour")
	}

	msg, err := c.parkingUsecase.FindRegisNoForCarByColour(args[0])
	if err != nil {
		return err
	}
	fmt.Println(msg)
	return nil
}

func (c *ParkingController) findSlotNumberByColour(args []string) error {
	if len(args) == 0 {
		return errors.New("Required colour")
	}

	msg, err := c.parkingUsecase.FindSlotNumberByColour(args[0])
	if err != nil {
		return err
	}
	fmt.Println(msg)
	return nil
}

func (c *ParkingController) findSlotNumberByRegisNo(args []string) error {
	if len(args) == 0 {
		return errors.New("Required registration No")
	}

	msg, err := c.parkingUsecase.FindSlotNumberByRegisNo(args[0])
	if err != nil {
		return err
	}
	fmt.Println(msg)
	return nil
}

func (c *ParkingController) ProcessCommand(command string) (matchedCommand bool, err error) {
	commands := strings.Fields(command)

	command = commands[0]
	var arrAgrs []string
	if len(commands) > 1 {
		arrAgrs = commands[1:]
	}

	switch command {
	case "exit":
		os.Exit(0)
	case "create_parking_lot":
		return true, c.createParkingLot(arrAgrs)
	case "park":
		return true, c.park(arrAgrs)
	case "leave":
		return true, c.leave(arrAgrs)
	case "status":
		return true, c.status(arrAgrs)
	case "registration_numbers_for_cars_with_colour":
		return true, c.findRegisNoForCarByColour(arrAgrs)
	case "slot_numbers_for_cars_with_colour":
		return true, c.findSlotNumberByColour(arrAgrs)
	case "slot_number_for_registration_number":
		return true, c.findSlotNumberByRegisNo(arrAgrs)
	}

	return false, nil
}

func NewParkingController(_parkingUsecase parking.Usecase) *ParkingController {
	return &ParkingController{parkingUsecase: _parkingUsecase}
}
