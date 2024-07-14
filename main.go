package main

import (
	. "ParkingLot/constants"
	. "ParkingLot/model"
	. "ParkingLot/service"
)

func main() {
	parkingLotService := &ParkingLotService{}
	parkingLotService.CreateParkingLot(2)

	car1 := Car{
		RegistrationNumber: "KA-01-HH-1234",
		Color:              White,
	}
	parkingLotService.Park(car1)
	parkingLotService.GetParkingLotStatus()

	car2 := Car{
		RegistrationNumber: "KA-01-HH-5678",
		Color:              Black,
	}
	parkingLotService.Park(car2)
	parkingLotService.GetParkingLotStatus()

	parkingLotService.Leave(0)
	parkingLotService.GetParkingLotStatus()

	parkingLotService.Leave(1)
	parkingLotService.GetParkingLotStatus()
}

/*
parking lot
lots []slots
capacity


Slot
availability
slot id
Car


Car
registration number
color

Color
const

ParkingLotService
CreateParkingLot
Park
Leave
GetParkingLotStatus

*/
