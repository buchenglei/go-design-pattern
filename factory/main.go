package main

import "fmt"

// v1 test
// func main() {
// 	pizzaStore := NewPizzaStore(SimplePizzaFactory{})

// 	pizzaStore.OrderPizza("cheese")
// 	pizzaStore.OrderPizza("clam")
// 	pizzaStore.OrderPizza("default")
// }

// v2 test
// func main() {
// 	nystore := NewNYStylePizzaStore()

// 	nystore.OrderPizza("cheese")
// 	fmt.Println("-----------")
// 	nystore.OrderPizza("clam")
// 	fmt.Println("-----------")
// 	nystore.OrderPizza("随便")

// 	fmt.Println("\n>>>>>>>>>>>\n")

// 	chicagostore := NewChicagoStylePizzaStore()

// 	chicagostore.OrderPizza("cheese")
// 	fmt.Println("-----------")
// 	chicagostore.OrderPizza("clam")
// 	fmt.Println("-----------")
// 	chicagostore.OrderPizza("随便")
// }

// v2.1 test
func main() {
	// 创建一家使用了纽约原料工厂的纽约pizza店
	pizzaStore := NewAnotherNYStylePizzaStore(NYPizzaIngredientFactory{})

	myPIzza := pizzaStore.OrderPizza("cheese")

	fmt.Println("\n==============\n")
	fmt.Printf("Final PIzza: %+v", myPIzza)
}
