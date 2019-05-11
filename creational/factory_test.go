package creational

import (
	"testing"
)

func TestBlackWidow_InfinityStone(t *testing.T) {
	avenger, err := GetAvenger(BWidow)
	if err != nil {
		t.Fatal("Avenger Black widow must exist but it is not ")
	}
	blackWidow := avenger.SetInfinityStone("Mind Stone")
	if blackWidow.infinityStone != "Mind Stone" {
		t.Error("Infinity Stone has not been set correctly")
	}
	t.Log("Log: black widow stone is -> ", blackWidow.infinityStone)
}

func TestCaptainAmerica_SuperPower(t *testing.T) {
	avenger, err := GetAvenger(CAmerica)
	if err != nil {
		t.Fatal("Avenger Captain America must exist but it is not ")
	}
	captainAmerica := avenger.SetSuperPower("Reality Stone")
	if captainAmerica.superPower != "Reality Stone" {
		t.Error("Infinity Stone has not been set correctly")
	}
	t.Log("Log: Cap  stone is -> ", captainAmerica.infinityStone)
}

func TestIronMan_SuperPower(t *testing.T) {
	avenger, err := GetAvenger(IMan)
	if err != nil {
		t.Fatal("Avenger Iron Man must exist but it is not ")
	}
	ironMan := avenger.SetInfinityStone("Power Stone")
	if ironMan.infinityStone != "Power Stone" {
		t.Error("Infinity Stone has not been set correctly")
	}
	t.Log("Log: Iron Man stone is -> ", ironMan.infinityStone)
}
