package structural

// first example
type Athlete struct{}

func (a *Athlete) Train() {
	println("Training")
}

func Swim() {
	println("Swimming")
}

// *func is like to be able to assign a complete different function to it
type CompositeSwimmerA struct {
	MyAthlete Athlete
	MySwim    func()
}

// second example with interfaces for more clear code

type Trainer interface {
	Train()
}

type Swimmer interface {
	Swim()
}
type SwimmerImplementor struct{}

func (s *SwimmerImplementor) Swim() {
	println("Swimming ... ")
}

type CompositeSwimmerB struct {
	Trainer
	Swimmer
}

// .......... embedded object example

type Animal struct{}

func (a *Animal) Eat() {
	println("Eating ... ")
}

// embedded objects make use of Zero initialization feature in GO
type Shark struct {
	Animal
	SharkSwim func()
}
