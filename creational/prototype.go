package creational

import (
	"errors"
	"fmt"
)

// simulate a t-shirt shop
type ShirtCloner interface {
	GetClone(m int) (ItemInfoGetter, error)
}

const (
	White = 1
	Black = 2
	Blue  = 3
)

type ShirtsCashe struct{}

func (s *ShirtsCashe) GetClone(m int) (ItemInfoGetter, error) {
	switch m {
	case White:
		newItem := *whitePrototype
		return &newItem, nil
	case Black:
		newItem := *blackPrototype
		return &newItem, nil
	case Blue:
		newItem := *bluePrototype
		return &newItem, nil
	}
	return nil, errors.New("shirt model is not recognized")
}

type ItemInfoGetter interface {
	GetInfo() string
}
type ShirtColor byte
type Shirt struct {
	Price float32
	SKU   string
	Color ShirtColor
}

var whitePrototype *Shirt = &Shirt{
	Price: 20.00,
	SKU:   "empty",
	Color: White,
}

var blackPrototype *Shirt = &Shirt{
	Price: 53.00,
	SKU:   "empty",
	Color: Black,
}
var bluePrototype *Shirt = &Shirt{
	Price: 62.00,
	SKU:   "empty",
	Color: Blue,
}

func (s *Shirt) GetInfo() string {
	return fmt.Sprintf("Shirt with SKU %s and color %d that costs %f\n", s.SKU, s.Color, s.Price)
}

func GetShirtsCloner() ShirtCloner {
	return new(ShirtsCashe)
}

func (s *Shirt) GetPrice() float32 {
	return s.Price
}
