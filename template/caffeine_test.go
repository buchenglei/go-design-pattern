package template

import (
	"testing"
	"fmt"
)

func TestCoffeine(t *testing.T) {
	tea := Tea{}
	coffee := Coffee{}

	testFunc := func(coffeine ProduceCoffeine) {
		coffeine.BoilWater()
		coffeine.Brew()
		coffeine.AddCondiments()
		coffeine.PourInCup()
	}

	testFunc(tea)
	fmt.Println("=======================")
	testFunc(coffee)
}