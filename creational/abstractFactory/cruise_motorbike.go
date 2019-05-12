package abstractFactory

type CruiseMotorbike struct{}

func (l *CruiseMotorbike) GetDoors() int {
	return 0
}

func (l *CruiseMotorbike) GetWheels() int {
	return 2
}

func (l *CruiseMotorbike) GetSeats() int {
	return 2
}
