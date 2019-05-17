package structural

import (
	"fmt"
	"testing"
)

func TestCar_Manufacture(t *testing.T) {
	car := Car{}
	assemble := &Assemble{}
	produce := &Produce{}

	msg, err := car.Manufacture(assemble, produce)
	if err != nil {
		t.Errorf("Error trying to use the car implementation: meassage: %s", err.Error())
	}
	fmt.Printf("%s\n", msg)
}

func TestBus_Manufacture(t *testing.T) {
	bus := Bus{}
	assemble := &Assemble{}
	produce := &Produce{}

	msg, err := bus.Manufacture(assemble, produce)
	if err != nil {
		t.Errorf("Error trying to use the car implementation: meassage: %s", err.Error())
	}
	fmt.Printf("%s\n", msg)
}
