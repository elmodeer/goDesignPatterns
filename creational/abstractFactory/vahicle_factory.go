package abstractFactory

import "errors"

type VehicleFactory interface {
	NewVehicle(v int) (Vehicle, error)
	GetVehicle(v int) (Vehicle, error)
}

const (
	CarFactoryType       = 1
	MotorbikeFactoryType = 2
)

func GetVehicleFactory(f int) (VehicleFactory, error) {
	switch f {
	case CarFactoryType:
		return new(CarFactory), nil
	case MotorbikeFactoryType:
		return new(MotorbikeFactory), nil
	}
	return nil, errors.New("unrecognized factory type")

}
