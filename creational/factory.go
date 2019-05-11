package creational

import (
	"errors"
	"fmt"
)

type AvengersInterface interface {
	SetSuperPower(s string) Avenger
	SetInfinityStone(is string) Avenger
}

const (
	IMan     = 1
	CAmerica = 2
	BWidow   = 3
	Hk       = 4
)

func GetAvenger(m int) (AvengersInterface, error) {
	switch m {
	case IMan:
		return new(IronMan), nil
	case CAmerica:
		return new(CaptainAmerica), nil
	case BWidow:
		return new(BlackWidow), nil
	}
	return nil, errors.New(fmt.Sprint(m, "Avenger type -> %s is not recognized"))
}

type Avenger struct {
	superPower    string
	infinityStone string
}

// Iron Man
type IronMan struct {
	a Avenger
}

func (im *IronMan) SetSuperPower(s string) Avenger {
	im.a.superPower = s
	return im.a
}

func (im *IronMan) SetInfinityStone(is string) Avenger {
	im.a.infinityStone = is
	return im.a
}

// Captain America
type CaptainAmerica struct {
	a Avenger
}

func (c *CaptainAmerica) SetSuperPower(s string) Avenger {
	c.a.superPower = s
	return c.a
}

func (c *CaptainAmerica) SetInfinityStone(is string) Avenger {
	c.a.infinityStone = is
	return c.a
}

// black widow
type BlackWidow struct {
	a Avenger
}

func (b *BlackWidow) SetSuperPower(s string) Avenger {
	b.a.superPower = s
	return b.a
}

func (b *BlackWidow) SetInfinityStone(is string) Avenger {
	b.a.infinityStone = is
	return b.a
}
