package parking

type Usecase interface {
	CreateParkLot(num uint16) (string, error)
	Park(regisNo string, colour string) (string, error)
	Leave(slot uint16) (string, error)
	FindRegisNoForCarByColour(colour string) (string, error)
	FindSlotNumberByColour(colour string) (string, error)
	FindSlotNumberByRegisNo(regisNo string) (string, error)
	Status() (string, error)
}

type Repository interface {
	CreateParkLot(num uint16) error
	Park(regisNo string, colour string) (uint16, error)
	Leave(slot uint16) error
	List(regisNo string, colour string) ([]Slot, error)
}

type Car struct {
	RegisNo string
	Colour  string
}

type Slot struct {
	Car     *Car
	SlotNo  uint16
	IsEmpty bool
}

type ParkingLot struct {
	Slots      []Slot
	EmptySlots uint16
}
