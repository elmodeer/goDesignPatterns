package abstractFactory

type LuxuryCar struct{}

func (l *LuxuryCar) GetDoors() int {
	return 6
}

func (l *LuxuryCar) GetWheels() int {
	return 8
}

func (l *LuxuryCar) GetSeats() int {
	return 9
}
