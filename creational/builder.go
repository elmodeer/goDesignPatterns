package creational

type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

// Director
// very good candidate to be of singleton
type ManufacturingDirector struct {
	builder BuildProcess
}

func (f *ManufacturingDirector) Construct() {
	f.builder.SetSeats().SetWheels().SetStructure()
}

func (f *ManufacturingDirector) SetBuilder(b BuildProcess) {
	f.builder = b
}

type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

// car Builder
type CarBuilder struct {
	v VehicleProduct
}

func (c *CarBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 4
	return c
}

func (c *CarBuilder) SetSeats() BuildProcess {
	c.v.Seats = 4
	return c
}
func (c *CarBuilder) SetStructure() BuildProcess {
	c.v.Structure = "car"
	return c
}

func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.v
}

// bike Builder
type BikeBuilder struct {
	v VehicleProduct
}

func (b *BikeBuilder) SetWheels() BuildProcess {
	b.v.Wheels = 2
	return b
}

func (b *BikeBuilder) SetSeats() BuildProcess {
	b.v.Seats = 1
	return b
}
func (b *BikeBuilder) SetStructure() BuildProcess {
	b.v.Structure = "bike"
	return b
}

func (b *BikeBuilder) GetVehicle() VehicleProduct {
	return b.v
}

// bus Builder
type BusBuilder struct {
	v VehicleProduct
}

func (bu *BusBuilder) SetWheels() BuildProcess {
	bu.v.Wheels = 6
	return bu
}

func (bu *BusBuilder) SetSeats() BuildProcess {
	bu.v.Seats = 16
	return bu
}
func (bu *BusBuilder) SetStructure() BuildProcess {
	bu.v.Structure = "bus"
	return bu
}

func (bu *BusBuilder) GetVehicle() VehicleProduct {
	return bu.v
}
