package main

import "fmt"

/********** 定义Pizza要实现的方法 **********/
type Pizzable interface {
	Prepare()
	Bake()
	Box()
	Cut()
}

// Pizza 将pizza定义为一个类，它本身可以包含一些有用的实现
// 这些实现可以被覆盖
type Pizza struct {
	Name string
}

// 这里是pizza的一些默认方法
func (pizza Pizza) Prepare() {
	fmt.Printf("Prepare %s pizza\n", pizza.Name)
}
func (pizza Pizza) Bake() {
	fmt.Printf("Bake %s pizza\n", pizza.Name)
}
func (pizza Pizza) Cut() {
	fmt.Printf("Cut %s pizza\n", pizza.Name)
}
func (pizza Pizza) Box() {
	fmt.Printf("Box %s pizza\n", pizza.Name)
}

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
