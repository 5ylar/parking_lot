package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	parkingController "parking_lot/parking/controller"
	parkingRepository "parking_lot/parking/repository"
	parkingUsecase "parking_lot/parking/usecase"
)

func main() {
	parkingRepo := parkingRepository.NewParkingRepository()
	parkingUsecase := parkingUsecase.NewParkingUsecase(parkingRepo)
	_parkingController := parkingController.NewParkingController(parkingUsecase)

	args := os.Args[1:]
	if len(args) > 0 {
		file, err := os.Open(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			command := scanner.Text()
			matched, err := _parkingController.ProcessCommand(command)
			if err != nil {
				fmt.Println(err.Error())
			}
			if matched {
				continue
			}
			/* Can use other controllers */
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

	} else {
		reader := bufio.NewReader(os.Stdin)
		for {
			command, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println(err)
			}
			matched, err := _parkingController.ProcessCommand(command)
			if err != nil {
				fmt.Println(err.Error())
			}
			if matched {
				continue
			}
			/* Can use other controllers */
		}
	}
}
