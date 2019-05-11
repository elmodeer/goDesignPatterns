package designPatterns

import "testing"

func TestGetInstance(t *testing.T) {
	counter1 := GetInstance()
	if counter1 == nil {
		// test of acceptance criteria 1 failed
		t.Error("Expected pointer to Singleton after calling GetInstance(), not nil ")
	}
	expectedCounter := counter1
	currentCount := counter1.AddOne()
	if currentCount != 1 {
		t.Errorf("Aftet calling for the first time to count, the count must be 1 but it is %d\n", currentCount)
	}
	counter2 := GetInstance()
	if counter2 != expectedCounter {
		t.Error("Expected same instance in counter2 but it got a different instance")
	}
	currentCount = counter2.AddOne()
	if currentCount != 2 {
		t.Errorf("After calling addOne second time the current value must be 2 but it is %d\n", currentCount)
	}
}
