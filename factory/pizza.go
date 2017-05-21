package main

import "fmt"

/********** 定义一个全新的Pizza要实现的方法 **********/
type Pizzable interface {
	Prepare()
	Bake()
	Box()
	Cut()
}

// Pizza 将pizza定义为一个类，它本身可以包含一些有用的实现
// 这些实现可以被覆盖
type Pizza struct {
	Name     string
	Dough    string
	Sauce    string
	Toppings []string
}

// 这里是pizza的一些默认方法
func (pizza Pizza) Prepare() {
	fmt.Println("Pareparing " + pizza.Name)
	fmt.Println("\t>Tossing dough...")
	fmt.Println("\t>Adding sauce...")
	fmt.Println("\t>Adding toppings:")
	for _, topping := range pizza.Toppings {
		fmt.Println("\t\t>" + topping)
	}
}
func (pizza Pizza) Bake() {
	fmt.Println("  >Bake for 25 minutes at 350")
}
func (pizza Pizza) Cut() {
	fmt.Println("  >Cut the pizza into diagonal slices")
}
func (pizza Pizza) Box() {
	fmt.Println("  >Place pizza in official PizzaStore box")
}
