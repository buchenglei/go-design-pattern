package main

import "fmt"

type Beverage interface {
	Description() string
	Cost() float32
}

/********** 基本的饮料类型 **********/

// Espresso 浓缩咖啡的数据结构
type Espresso struct {
	description string
}

func NewEspresso() *Espresso {
	return &Espresso{
		description: "Espresso",
	}
}

func (espresso Espresso) Description() string {
	return espresso.description
}

func (Espresso) Cost() float32 {
	return 1.99
}

// HouseBlend 另一种饮料的数据结构
type HouseBlend struct {
	description string
}

func NewHouseBlend() *HouseBlend {
	return &HouseBlend{
		description: "HouseBlend",
	}
}

func (houseBlend HouseBlend) Description() string {
	return houseBlend.description
}

func (HouseBlend) Cost() float32 {
	return 0.89
}

/********** 调料对象 **********/

// Mocha 摩卡调料
type Mocha struct {
	beverage Beverage // 保存装饰的对象
}

func NewMocha(b Beverage) *Mocha {
	return &Mocha{
		beverage: b,
	}
}

func (mocha Mocha) Description() string {
	return mocha.beverage.Description() + ", Mocha"
}

func (mocha Mocha) Cost() float32 {
	return 0.20 + mocha.beverage.Cost()
}

// Milk 牛奶调料
type Milk struct {
	beverage Beverage // 保存装饰的对象
}

func NewMilk(b Beverage) *Milk {
	return &Milk{
		beverage: b,
	}
}

func (milk Milk) Description() string {
	return milk.beverage.Description() + ", Milk"
}

func (milk Milk) Cost() float32 {
	return 0.20 + milk.beverage.Cost()
}

func main() {
	// 浓缩咖啡饮料
	espresso := NewEspresso()
	// 综合咖啡饮料
	houseBlend := NewHouseBlend()

	// 摩卡浓缩
	mochaEspresso := NewMocha(espresso)
	// 摩卡综合
	mochahouseBlend := NewMocha(houseBlend)

	// 牛奶浓缩
	milkEspresso := NewMilk(espresso)
	// 牛奶综合
	milkHouseBlend := NewMilk(houseBlend)

	// 双倍牛奶浓缩
	doubleMilkEspresso := NewMilk(NewMilk(espresso))

	fmt.Printf("%s 售价: ￥%.2f\n", mochaEspresso.Description(), mochaEspresso.Cost())
	fmt.Printf("%s 售价: ￥%.2f\n", mochahouseBlend.Description(), mochahouseBlend.Cost())
	fmt.Printf("%s 售价: ￥%.2f\n", milkEspresso.Description(), milkEspresso.Cost())
	fmt.Printf("%s 售价: ￥%.2f\n", milkHouseBlend.Description(), milkHouseBlend.Cost())
	fmt.Printf("%s 售价: ￥%.2f\n", doubleMilkEspresso.Description(), doubleMilkEspresso.Cost())
}
