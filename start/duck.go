package main

import "fmt"

// 定义不同类型的鸭子的行为
type (
	FlyBehavior interface {
		Fly()
	}

	QuackBehavior interface {
		Quack()
	}
)

// Duck class
type Duck struct {
	name          string
	flyBehavior   FlyBehavior
	quackBehavior QuackBehavior
}

// NewDuck 创建一只新的鸭子
func NewDuck(name string, Fly FlyBehavior, quack QuackBehavior) Duck {
	return Duck{
		name:          name,
		flyBehavior:   Fly,
		quackBehavior: quack,
	}
}

// Display 示例所有的鸭子共有的
func (duck Duck) Display() {
	fmt.Println("Hello, My name is", duck.name)
}

// Fly 定义鸭子的飞行行为
// 不同的鸭子飞行的方式不一样，故这里用不同的行为类去实现
func (duck Duck) Fly() {
	duck.flyBehavior.Fly()
}

// SetFlyBehavior 动态的设置鸭子的飞行行为
func (duck *Duck) SetFlyBehavior(flyBehavior FlyBehavior) {
	duck.flyBehavior = flyBehavior
}

// Quack 定义鸭子的叫的行为
// 不同的鸭子叫的方式不一样，故这里用不同的行为类去实现
func (duck Duck) Quack() {
	duck.quackBehavior.Quack()
}

// SetQuackBehavior 动态的设置鸭子的叫的行为
func (duck *Duck) SetQuackBehavior(quackBehavior QuackBehavior) {
	duck.quackBehavior = quackBehavior
}

// behavior classes
type (
	// fly behavior
	FlyWithWings struct{}
	FlyNoWay     struct{}

	// quack behavior
	Quack     struct{}
	MuteQuack struct{}
)

func (fly FlyWithWings) Fly() {
	fmt.Println("I'm flying with wings")
}

func (fly FlyNoWay) Fly() {
	fmt.Println("I can't fly")
}

func (quack Quack) Quack() {
	fmt.Println("Quack!")
}

func (quack MuteQuack) Quack() {
	fmt.Println("<< Slience >>")
}

func main() {
	// 一只会飞会叫的鸭子
	duckA := NewDuck(
		"duckA",
		FlyWithWings{},
		Quack{},
	)

	// 一只不会飞也不会叫的鸭子
	duckB := NewDuck(
		"duckB",
		FlyNoWay{},
		MuteQuack{},
	)

	// 通过这种方式可以组合出各种各样的鸭子

	duckA.Display()
	duckA.Fly()
	duckA.Quack()

	duckB.Display()
	duckB.Fly()
	duckB.Quack()

	// 现在可以任意改变鸭子的飞行和叫的行为
	duckA.SetFlyBehavior(FlyNoWay{})
	duckB.SetQuackBehavior(Quack{})

	duckA.Display()
	duckA.Fly()
	duckA.Quack()

	duckB.Display()
	duckB.Fly()
	duckB.Quack()
}
