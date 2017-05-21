package main

// 这是v2的一个小的升级

type AnotherNYStylePizzaStore struct {
	PizzaStoreV2

	ingredientFactory PizzaIngredientFactory // 定义这家店用到的原料工厂
}

func (nystore *AnotherNYStylePizzaStore) CreatePizza(name string) Pizzable {

	var pizza Pizzable

	switch name {
	case "cheese":
		// 纽约风格的店做出来的CheesePizza就需要用到纽约的原料工厂
		pizza = NewNewCheesePizza("NYStyle Cheese Pizza", nystore.ingredientFactory)
	}

	return pizza
}

func NewAnotherNYStylePizzaStore(factory PizzaIngredientFactory) PizzaSaler {
	nystore := &AnotherNYStylePizzaStore{}
	nystore.creater = nystore
	nystore.ingredientFactory = factory

	return nystore
}
