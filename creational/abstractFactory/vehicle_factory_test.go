package abstractFactory

import "testing"

func TestMotorbikeFactory(t *testing.T) {
	motorbikeFactory, err := GetVehicleFactory(MotorbikeFactoryType)
	if err != nil {
		t.Fatal(err)
	}
	sportMotorbike, err := motorbikeFactory.GetVehicle(2)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Motorbike vehicle has %d wheels and %d seats\n", sportMotorbike.GetWheels(), sportMotorbike.GetSeats())

	// test assertion
	//sporBike, ok := sportMotorbike.(Motorbike)
	//if !ok {
	//	t.Fatal("Struct assertion has failed")
	//}
	//t.Logf("sporbike has type %d\n", sporBike.GetType())
}

func TestCarFactory(t *testing.T) {
	carF, err := GetVehicleFactory(CarFactoryType)
	if err != nil {
		t.Fatal(err)
	}
	familyCar, err := carF.GetVehicle(2)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Motorbike vehicle has %d wheels and %d seats\n", familyCar.GetWheels(), familyCar.GetSeats())

	// test assertion
	//familyC, ok := familyCar.(Car)
	//if !ok {
	//	t.Fatal("Struct assertion has failed")
	//}
	//t.Logf("sporbike has type %d\n", familyC.GetType())
}
