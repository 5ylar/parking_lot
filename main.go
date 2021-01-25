package main

import (
	"bufio"
	"fmt"
	"os"
	// "os/exec"
	"strings"
	"strconv"
	"errors"
)

type ParkingLot struct {
	RegisNo string
	Colour string
}

var parkingLotsNum uint16
var parkingLots map[uint16]*ParkingLot

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		err = runCommand(cmdString)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
func runCommand(command string) error {
	commands := strings.Fields(command)

	command = commands[0]
	var arrAgrs []string
	if (len(commands) > 1) {
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

		parkingLots = make(map[uint16]*ParkingLot)
		parkingLotsNum = uint16(nLot)

		// for slot := 1; slot <= nLot; slot++ {
		// 	parkingLots[uint16(slot)] = nil
		// }

		fmt.Printf("Created a parking lot with %d slots\n",nLot)
		fmt.Println(parkingLots)
		return nil
	case "park":
		if len(arrAgrs) < 2 {
			return errors.New("Required at least 2 parameters (registeration number, colour)")
		}

		parkingLotsCurrentNum := uint16(len(parkingLots))
		var parkingLotsCursor uint16
		for slot := uint16(1); slot <= parkingLotsCurrentNum; slot++ {
			if parkingLots[uint16(slot)] == nil {
				parkingLotsCursor = slot
				break
			}
		}

		/* Parking lot is full (Last parking lot equal number of parking lots and not have any empty lot) */
		if (parkingLotsCurrentNum >= parkingLotsNum && parkingLotsCursor == 0) {
			return errors.New("Sorry, parking lot is full")
		}

		/* not have any empty lot and parking lot not full, set cursor to next lot */
		if parkingLotsCursor == 0 {
			parkingLotsCursor = parkingLotsCurrentNum + 1
		}

		/* insert lot */
		parkingLot := ParkingLot{
			RegisNo: arrAgrs[0],
			Colour: arrAgrs[1],
		}
		parkingLots[parkingLotsCursor] = &parkingLot
		
		return nil
	case "leave":
		if len(arrAgrs) == 0 {
			return errors.New("Required number of parking lot")
		}
		nLot, err := strconv.Atoi(arrAgrs[0])
		if err != nil {
			return errors.New("Wrong argument for parking lot")
		}

		parkingLots[uint16(nLot)] = nil
		return nil
	case "status":
		fmt.Println(parkingLots)
		return nil
	}

	return nil
	// cmd := exec.Command(arrCommandStr[0], arrCommandStr[1:]...)
	// cmd.Stderr = os.Stderr
	// cmd.Stdout = os.Stdout
	// return cmd.Run()
}