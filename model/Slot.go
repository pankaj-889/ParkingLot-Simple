package model

import "log"

type Slot struct {
	Availability bool `json:"availability,true"`
	SlotID       int  `json:"slot_id,omitempty"`
	Car          Car  `json:"car,omitempty"`
}

func NewSlot(slotID int) Slot {
	return Slot{
		Availability: true,
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
