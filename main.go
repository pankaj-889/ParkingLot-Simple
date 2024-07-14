package main

import (
	. "ParkingLot/constants"
	"log"
)

type ParkingLot struct {
	Slots    []Slot
	Capacity int
}

type IParkingLot interface {
	CreateParkingLot(capacity int)
	Park(car Car)
	Leave(slotNumber int)
	GetParkingLotStatus()
	GetNextSlot() int
}

func (p *ParkingLot) CreateParkingLot(capacity int) {
	p.Capacity = capacity
	p.Slots = make([]Slot, capacity)
	for i := range p.Slots {
		p.Slots[i] = NewSlot(i + 1)
	}
}

func (p *ParkingLot) Park(car Car) {
	slotNumber := p.GetNextSlot()
	p.Slots[slotNumber].Assign(car)
}

func (p *ParkingLot) Leave(slotNumber int) {
	p.Slots[slotNumber].UnAssign()
}

func (p *ParkingLot) GetParkingLotStatus() {
	for _, slot := range p.Slots {
		if slot.Availability {
			continue
		}
		car := slot.Car
		println("Slot Number: ", slot.SlotID)
		println("Registration Number: ", car.RegistrationNumber)
		println("Color: ", car.Color.String())
	}
}

func (p *ParkingLot) GetNextSlot() int {
	for i, slot := range p.Slots {
		if slot.Availability {
			return i
		}
	}
	return -1
}

type Slot struct {
	Availability bool `json:"availability,true"`
	SlotID       int  `json:"slot_id,omitempty"`
	Car          Car  `json:"car,omitempty"`
}

func NewSlot(slotID int) Slot {
	return Slot{
		Availability: true, // Set the default value to true
		SlotID:       slotID,
		Car:          Car{},
	}
}

type ISlot interface {
	Assign(car Car)
	UnAssign(car Car)
}

func (s *Slot) Assign(car Car) {
	if s.Availability {
		s.Car = car
		s.Availability = false
	}
}

func (s *Slot) UnAssign() {
	if s.Availability {
		log.Println("Slot is already empty")
	}
	s.Car = Car{}
	s.Availability = true
}

type Car struct {
	RegistrationNumber string
	Color              Color
}

func main() {
	parkingLot := ParkingLot{}
	parkingLot.CreateParkingLot(2)
	car := Car{
		RegistrationNumber: "KA-01-HH-1234",
		Color:              White,
	}
	parkingLot.Park(car)
	parkingLot.GetParkingLotStatus()
	parkingLot.Leave(0)
	parkingLot.GetParkingLotStatus()
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
