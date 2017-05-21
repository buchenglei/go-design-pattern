package main

type PizzaStoreV3 struct {
	creater PizzaCreater // 这是一个工厂方法
}

func (store PizzaStoreV3) OrderPizza(pizzaType string) Pizzable {
	pizza := store.creater.CreatePizza(pizzaType)

	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()

	return pizza
}

type AnotherNYStylePizzaStore struct {
	PizzaStoreV3

	ingredientFactory PizzaIngredientFactory // 定义这家店用到的原料工厂
}

func (nystore *AnotherNYStylePizzaStore) CreatePizza(name string) Pizzable {

	var pizza Pizzable

	switch name {
	case "cheese":
		// 纽约风格的店做出来的CheesePizza就需要用到纽约的原料工厂
		pizza = NewNewCheesePizza("Another NYStyle Cheese Pizza", nystore.ingredientFactory)
	}

	return pizza
}

func NewAnotherNYStylePizzaStore(factory PizzaIngredientFactory) PizzaSaler {
	nystore := &AnotherNYStylePizzaStore{}
	nystore.creater = nystore
	nystore.ingredientFactory = factory

	return nystore
}
