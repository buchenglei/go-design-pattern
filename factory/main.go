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
func main() {
	nystore := NewNYStylePizzaStore()

	nystore.OrderPizza("cheese")
	fmt.Println("-----------")
	nystore.OrderPizza("clam")
	fmt.Println("-----------")
	nystore.OrderPizza("随便")

	fmt.Println("\n>>>>>>>>>>>\n")

	chicagostore := NewChicagoStylePizzaStore()

	chicagostore.OrderPizza("cheese")
	fmt.Println("-----------")
	chicagostore.OrderPizza("clam")
	fmt.Println("-----------")
	chicagostore.OrderPizza("随便")
}
