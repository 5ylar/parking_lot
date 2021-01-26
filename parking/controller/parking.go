package controllor

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"parking_lot/parking"
	"strconv"
	"strings"
)

type ParkingController struct {
	parkingUsecase parking.Usecase
}

func InitParkingController(_parkingUsecase parking.Usecase) {
	args := os.Args[1:]
	if len(args) > 0 {
		file, err := os.Open(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			cmdString := scanner.Text()
			err := runCommand(cmdString, _parkingUsecase)
			if err != nil {
				fmt.Println(err.Error())
			}
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

	} else {
		reader := bufio.NewReader(os.Stdin)
		for {
			cmdString, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println(err)
			}
			err = runCommand(cmdString, _parkingUsecase)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}
}

func runCommand(command string, parkingUsecase parking.Usecase) error {
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
		if len(arrAgrs) == 0 {
			return errors.New("Required number of parking lot")
		}
		nLot, err := strconv.Atoi(arrAgrs[0])
		if err != nil {
			return errors.New("Wrong argument for parking lot")
		}

		msg, err := parkingUsecase.CreateParkLot(uint16(nLot))
		if err != nil {
			return err
		}
		fmt.Println(msg)

		return nil
	case "park":
		if len(arrAgrs) < 2 {
			return errors.New("Required at least 2 parameters (registeration number, colour)")
		}

		msg, err := parkingUsecase.Park(arrAgrs[0], arrAgrs[1])
		if err != nil {
			return err
		}
		fmt.Println(msg)

		return nil
	case "leave":
		if len(arrAgrs) == 0 {
			return errors.New("Required number of parking lot")
		}
		nLot, err := strconv.Atoi(arrAgrs[0])
		if err != nil {
			return errors.New("Wrong argument for parking lot")
		}

		msg, err := parkingUsecase.Leave(uint16(nLot))
		if err != nil {
			return err
		}
		fmt.Println(msg)
		return nil
	case "status":
		msg, err := parkingUsecase.Status()
		if err != nil {
			return err
		}
		fmt.Println(msg)
		return nil
	case "registration_numbers_for_cars_with_colour":
		if len(arrAgrs) == 0 {
			return errors.New("Required colour")
		}

		msg, err := parkingUsecase.FindRegisNoForCarByColour(arrAgrs[0])
		if err != nil {
			return err
		}
		fmt.Println(msg)
		return nil
	case "slot_numbers_for_cars_with_colour":
		if len(arrAgrs) == 0 {
			return errors.New("Required colour")
		}

		msg, err := parkingUsecase.FindSlotNumberByColour(arrAgrs[0])
		if err != nil {
			return err
		}
		fmt.Println(msg)
		return nil
	case "slot_number_for_registration_number":
		if len(arrAgrs) == 0 {
			return errors.New("Required registration No")
		}

		msg, err := parkingUsecase.FindSlotNumberByRegisNo(arrAgrs[0])
		if err != nil {
			return err
		}
		fmt.Println(msg)
		return nil
	}

	return nil
}
