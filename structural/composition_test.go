package structural

import "testing"

func TestSwimmer_Swim(t *testing.T) {
	swimmer := CompositeSwimmerA{
		MySwim: Swim,
	}
	swimmer.MyAthlete.Train()
	swimmer.MySwim()

}
func TestAnimal_Swim(t *testing.T) {
	fish := Shark{
		SharkSwim: Swim,
	}
	fish.Eat()
	fish.SharkSwim()
}

func TestSwimmerImplementor_Swim(t *testing.T) {
	swimmer := CompositeSwimmerB{
		&Athlete{},
		&SwimmerImplementor{},
	}
	swimmer.Swim()
	swimmer.Train()
}
