package abstractFactory

type SportMotorbike struct{}

func (l *SportMotorbike) GetDoors() int {
	return 0
}

func (l *SportMotorbike) GetWheels() int {
	return 2
}

func (l *SportMotorbike) GetSeats() int {
	return 1
}
