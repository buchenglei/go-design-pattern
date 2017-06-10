package template

import(
	"fmt"
)

// ProduceCoffeine 咖啡和茶都要实现的方法
type ProduceCoffeine interface {
	BoilWater()
	Brew()
	AddCondiments()
	PourInCup()
}

type (
	// 基类
	CaffeineBeverage struct {}

	Tea struct {
		CaffeineBeverage
	}

	Coffee struct {
		CaffeineBeverage
	}
)

func (CaffeineBeverage) BoilWater() {
	fmt.Println("Boiling water!")
}

func (CaffeineBeverage) PourInCup() {
	fmt.Println("Boiling water!")
}

func (Tea) Brew() {
	fmt.Println("Steeping the tea")
}

func (Tea) AddCondiments() {
	fmt.Println("Pouring into tea")
}

func (Coffee) Brew() {
	fmt.Println("Steeping the coffee")
}

func (Coffee) AddCondiments() {
	fmt.Println("Pouring into coffee")
}


