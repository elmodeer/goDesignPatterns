package creational

import (
	"testing"
)

func TestClone(t *testing.T) {
	shirtsCashe := GetShirtsCloner()
	if shirtsCashe == nil {
		t.Fatal("Not implemented yet")
	}
	item1, err := shirtsCashe.GetClone(White)
	if err != nil {
		t.Error(err)
	}
	if item1 == whitePrototype {
		t.Fatal("items cant be equal, that means cloning is not working")
	}

	shirt1, ok := item1.(*Shirt)
	if !ok {
		t.Fatal("type assertion for shirt1 failed")
	}
	shirt1.SKU = "abbcc"
}
