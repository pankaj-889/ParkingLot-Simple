package Interfaces

import (
	. "ParkingLot/model"
)

type IParkingLotService interface {
	CreateParkingLot(lot int)
	GetParkingLotStatus()
	Park(car Car)
	Leave(slotNumber int)
}

type IParkingLot interface {
	CreateParkingLot(capacity int)
	Park(car Car)
	Leave(slotNumber int)
	GetParkingLotStatus()
	GetNextSlot() int
}
