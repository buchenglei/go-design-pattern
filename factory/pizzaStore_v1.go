package main

import "fmt"

/********** 一些其他的pizza **********/
type CheesePizza struct {
	Pizza
}

type ClamPizza struct {
	Pizza
}

/********** 简单工厂 ***********/
// 简单工厂并不是一种设计模式，而是一种编程习惯
type SimplePizzaFactory struct{}

func (SimplePizzaFactory) CreatePizza(name string) Pizzable {
	var pizza Pizzable
	switch name {
	case "cheese":
		pizza = &CheesePizza{
			Pizza: Pizza{Name: name},
		}
	case "clam":
		pizza = &CheesePizza{
			Pizza: Pizza{Name: name},
		}
	default:
		pizza = &Pizza{
			Name: "default",
		} // 默认情况返回一个默认的pizza
	}

	return pizza
}

type PizzaStore struct {
	factory SimplePizzaFactory
}

func NewPizzaStore(factory SimplePizzaFactory) *PizzaStore {
	return &PizzaStore{
		factory: factory,
	}
}

func (store PizzaStore) OrderPizza(name string) Pizzable {
	var pizza Pizzable

	pizza = store.factory.CreatePizza(name)

	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()

	fmt.Println("> pizza done")

	return pizza
}
