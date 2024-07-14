package model

type ParkingLot struct {
	Slots    []Slot
	Capacity int
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
