package Service

import (
	. "ParkingLot"
)

type ParkingLotService struct {
	ParkingLot ParkingLot
}

type IParkingLotService interface {
	CreateParkingLot(lot ParkingLot)
	GetParkingLot() ParkingLot
	Park(car Car)
	Leave(slotNumber int)
}

func (p *ParkingLotService) CreateParkingLot(lot ParkingLot) {
	p.CreateParkingLot(lot)
}

func (p *ParkingLotService) GetParkingLot() ParkingLot {
	return p.GetParkingLot()
}

func (p *ParkingLotService) Park(car Car) {
	p.Park(car)
}

func (p *ParkingLotService) Leave(slotNumber int) {
	p.Leave(slotNumber)
}
