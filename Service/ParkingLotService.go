package Service

import (
	. "ParkingLot/model"
)

type ParkingLotService struct {
	ParkingLot *ParkingLot
}

func (p *ParkingLotService) CreateParkingLot(lot int) {
	p.CreateParkingLot(lot)
}

func (p *ParkingLotService) GetParkingLotStatus() ParkingLot {
	return p.GetParkingLotStatus()
}

func (p *ParkingLotService) Park(car Car) {
	p.Park(car)
}

func (p *ParkingLotService) Leave(slotNumber int) {
	p.Leave(slotNumber)
}
