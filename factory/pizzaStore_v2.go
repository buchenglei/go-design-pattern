package main

import "fmt"

type (
	PizzaOrder interface {
		OrderPizza(string) Pizzable
	}

	PizzaCreater interface {
		CreatePizza(string) Pizzable
	}

	PizzaSaler interface {
		PizzaOrder
		PizzaCreater
	}
)

// 工厂方法模式 定义了一个创建对象的接口，但是由子类决定要实例化的类是哪一个
// 工厂方法让类把实例化的操作推迟到了子类中去

type PizzaStoreV2 struct {
	creater PizzaCreater // 这是一个工厂方法
}

func (store PizzaStoreV2) OrderPizza(pizzaType string) Pizzable {
	pizza := store.creater.CreatePizza(pizzaType)

	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()

	return pizza
}

/********** 纽约风格的pizza店铺 **********/

// NYStyleCheesePizza 纽约风格的pizza
type NYStyleCheesePizza struct {
	Pizza
}

// NYStyleClamPizza 纽约风格的pizza
type NYStyleClamPizza struct {
	Pizza
}

type NYStylePizzaStore struct {
	PizzaStoreV2
}

func (nystore *NYStylePizzaStore) CreatePizza(name string) Pizzable {

	var pizza Pizzable

	switch name {
	case "cheese":
		pizza = &NYStyleCheesePizza{
			Pizza: Pizza{
				Name: "NYStyle " + name,
				Toppings: []string{
					"Grated Raggiano Cheese",
				}, // 加一些纽约店的特设
			},
		}
	case "clam":
		pizza = &NYStyleClamPizza{
			Pizza: Pizza{Name: "NYStyle " + name},
		}
	default:
		pizza = &Pizza{
			Name: "NYStyle default",
		} // 默认情况返回一个默认的pizza
	}

	return pizza
}

func NewNYStylePizzaStore() PizzaSaler {
	nystore := &NYStylePizzaStore{}
	nystore.creater = nystore

	return nystore
}

/********** 芝加哥风格的pizza店 **********/

// ChicagoStyleCheesePizza 芝加哥风格的pizza
type ChicagoStyleCheesePizza struct {
	Pizza
}

func (pizza ChicagoStyleCheesePizza) Cut() {
	fmt.Println("  >Chicago Style Cheese Pizza will be cutted into square slices")
}

// ChicagoStyleClamPizza 芝加哥风格的pizza
type ChicagoStyleClamPizza struct {
	Pizza
}

type ChicagoStylePizzaStore struct {
	PizzaStoreV2
}

func (nystore *ChicagoStylePizzaStore) CreatePizza(name string) Pizzable {

	var pizza Pizzable

	switch name {
	case "cheese":
		pizza = &ChicagoStyleCheesePizza{
			Pizza: Pizza{Name: "ChicagoStyle " + name},
		}
	case "clam":
		pizza = &ChicagoStyleClamPizza{
			Pizza: Pizza{Name: "ChicagoStyle " + name},
		}
	default:
		pizza = &Pizza{
			Name: "ChicagoStyle default",
		} // 默认情况返回一个默认的pizza
	}

	return pizza
}

func NewChicagoStylePizzaStore() PizzaSaler {
	chicagostore := &ChicagoStylePizzaStore{}
	chicagostore.creater = chicagostore

	return chicagostore
}
