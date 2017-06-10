package template

import (
	"fmt"
)

// CoffeineDrink 实际的饮料需要实现的部分方法
type CoffeineDrink interface {
	Brew()
	AddCondiments()
	// 这个方法就是一个Hook方法
	WantCondiments() bool 
}

type Preparer interface {
	PrepareRecipe()
}

// CaffeineBeverageWithHook
type CaffeineBeverageWithHook struct {
	Drink CoffeineDrink
}

func (coffee CaffeineBeverageWithHook) PrepareRecipe() {
	coffee.Drink.Brew()
	coffee.boilWater()
	coffee.pourInCup()

	if coffee.Drink.WantCondiments() {
		coffee.Drink.AddCondiments()
	}
}

func (CaffeineBeverageWithHook) boilWater() {
	fmt.Println("Boiling water!")
}

func (CaffeineBeverageWithHook) pourInCup() {
	fmt.Println("Pouring into cup!")
}

type CoffeeWithHook struct {
	CaffeineBeverageWithHook
}

func NewCoffeeWithHook() *CoffeeWithHook {
	coffee := &CoffeeWithHook{}

	coffee.Drink = coffee

	return coffee
}

func (CoffeeWithHook) Brew() {
	fmt.Println("Steeping the coffee")
}

func (CoffeeWithHook) AddCondiments() {
	fmt.Println("Add condiments into coffee")
}

func (CoffeeWithHook) WantCondiments() bool {
	return false
}
