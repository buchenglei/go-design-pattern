package main

import "fmt"

type PizzaPreparer interface {
	Prepare()
}

type NewPizza struct {
	Name      string
	Dough     string
	Sauce     string
	Vaggies   []string
	Cheese    string
	Pepperoni string
	Clam      string

	Preparer PizzaPreparer
}

func (pizza NewPizza) Bake() {
	fmt.Println("  >Bake for 25 minutes at 350")
}
func (pizza NewPizza) Cut() {
	fmt.Println("  >Cut the pizza into diagonal slices")
}
func (pizza NewPizza) Box() {
	fmt.Println("  >Place pizza in official PizzaStore box")
}

type NewCheesePizza struct {
	ingredientFactory PizzaIngredientFactory
	NewPizza
}

// NewNewCheesePizza 这里是创建一个新的CheesePizza
// 这是一个通用的方法，如果传进来的是纽约的原料工厂，
// 那么生产出来的就是纽约风味的CheesePizza
// 如果传进来的是芝加哥的原料工厂
// 那么生产出来的就是芝加哥风味的CheesePizza
func NewNewCheesePizza(name string, factory PizzaIngredientFactory) Pizzable {
	pizza := &NewCheesePizza{
		ingredientFactory: factory,
		NewPizza:          NewPizza{Name: name},
	}

	pizza.Preparer = pizza

	return pizza
}

func (pizza *NewCheesePizza) Prepare() {
	fmt.Println("Pareparing " + pizza.Name)
	pizza.Dough = pizza.ingredientFactory.CreateDough()
	pizza.Sauce = pizza.ingredientFactory.CreateSauce()
	pizza.Cheese = pizza.ingredientFactory.CreateCheese()
}

// 剩下的新的纽约的比萨店就可以参考pizzaStore_v2来设计
