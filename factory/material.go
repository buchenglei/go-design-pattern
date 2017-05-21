package main

// PizzaIngredientFactory 一个原料工厂需要实现的方法
type PizzaIngredientFactory interface {
	CreateDough() string
	CreateSauce() string
	CreateCheese() string
	CreateVeggies() []string
	CreatePepperoni() string
	CreateClam() string
}

// 这里的原料都用字符串表示，仅作示例

// 这个是纽约原料工程

type NYPizzaIngredientFactory struct{}

func (factory NYPizzaIngredientFactory) CreateDough() string {
	return "ThinCrustDough"
}

func (factory NYPizzaIngredientFactory) CreateSauce() string {
	return "MarinaraSauce"
}

func (factory NYPizzaIngredientFactory) CreateCheese() string {
	return "RegguabiCheese"
}

func (factory NYPizzaIngredientFactory) CreateVeggies() []string {
	return []string{
		"Garlic",
		"Onion",
		"Marshroom",
		"RedPapper",
	}
}

func (factory NYPizzaIngredientFactory) CreatePepperoni() string {
	return "SlicedPepperoni"
}

func (factory NYPizzaIngredientFactory) CreateClam() string {
	return "FreshClam"
}
