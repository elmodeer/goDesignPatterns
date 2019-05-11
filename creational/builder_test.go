package creational

import (
	"testing"
)

func TestBuilderPattern(t *testing.T) {
	manufacturingComplex := ManufacturingDirector{}

	// for car
	carBuilder := &CarBuilder{}
	manufacturingComplex.SetBuilder(carBuilder)
	manufacturingComplex.Construct()

	car := carBuilder.GetVehicle()
	//car.Seats = 2
	if car.Wheels != 4 {
		t.Errorf("Wheels are wrong, they should be 4 but it is %d\n", car.Wheels)
	}
	if car.Structure != "car" {
		t.Errorf("structure is wrong, it should be cat but it is %s\n", car.Structure)
	}
	if car.Seats != 4 {
		t.Errorf("seats are wrong, they should be 4 but it is %d\n", car.Seats)
	}

	// for bus
	busBuilder := &BusBuilder{}
	manufacturingComplex.SetBuilder(busBuilder)
	manufacturingComplex.Construct()

	bus := carBuilder.GetVehicle()
	if bus.Wheels != 4 {
		t.Errorf("Wheels are wrong, they should be 6 but it is %d\n", bus.Wheels)
	}
	if bus.Structure != "car" {
		t.Errorf("structure is wrong, it should be bus but it is %s\n", bus.Structure)
	}
	if bus.Seats != 4 {
		t.Errorf("seats are wrong, they should be 16 but it is %d\n", bus.Seats)
	}

}
