package Service

import (
	. "ParkingLot/model"
)

type ParkingLotService struct {
	ParkingLot ParkingLot
}

func (p *ParkingLotService) CreateParkingLot(lot int) {
	p.ParkingLot.CreateParkingLot(lot)
}

func (p *ParkingLotService) GetParkingLotStatus() {
	p.ParkingLot.GetParkingLotStatus()
}

func (p *ParkingLotService) Park(car Car) {
	p.ParkingLot.Park(car)
}

func (p *ParkingLotService) Leave(slotNumber int) {
	p.ParkingLot.Leave(slotNumber)
}
