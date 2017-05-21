package main

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

type PizzaStoreV2 struct {
	creater PizzaCreater
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
			Pizza: Pizza{Name: "NYStyle " + name},
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
