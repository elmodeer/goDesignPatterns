package structural

type Vehicle interface {
	Manufacture(Workshop, Workshop) (string, error)
}

type Car struct{}

func (c *Car) Manufacture(workshop1 Workshop, workshop2 Workshop) (string, error) {
	msgResult := "car is"
	if workshop1 != nil {
		msgResult += workshop1.Work()
	}
	if workshop2 != nil {
		msgResult += " and "
		msgResult += workshop2.Work()
	}
	return msgResult, nil
}

type Bus struct{}

func (c *Bus) Manufacture(workshop1 Workshop, workshop2 Workshop) (string, error) {
	msgResult := "bus is"
	if workshop1 != nil {
		msgResult += workshop1.Work()
	}
	if workshop2 != nil {
		msgResult += " and "
		msgResult += workshop2.Work()
	}
	return msgResult, nil
}

type Workshop interface {
	Work() string
}

type Assemble struct{}

func (a *Assemble) Work() string {
	msg := " assembled"
	return msg
}

type Produce struct{}

func (a *Produce) Work() string {
	msg := " produced"
	return msg
}
