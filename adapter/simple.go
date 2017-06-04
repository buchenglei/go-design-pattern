package adapter

import "fmt"

type (
	// Duck 定义鸭子的行为
	Duck interface {
		Quack()
		Fly()
	}

	// Turkey 火鸡的行为
	Turkey interface {
		Gobble()
		Fly()
	}
)

/****** 定义一些实现接口的具体实例 *****/

// MallarDuck 绿头鸭
type MallarDuck struct{}

func (MallarDuck) Quack() {
	fmt.Println("MarllarDuck Quack....")
}

func (MallarDuck) Fly() {
	fmt.Println("MarllarDuck Fly....")
}

type WildTurkey struct{}

func (WildTurkey) Gobble() {
	fmt.Println("WildTurkey Gobble....")
}

func (WildTurkey) Fly() {
	fmt.Println("WildTurkey fly a short distance....")
}

/***** 定义一个适配器 *****/

// TurkeyAdapter 火鸡适配器用来适配Duck接口
// 使火鸡看起来像一只鸭子
type TurkeyAdapter struct {
	turkey Turkey
}

func NewTurkeyAdapter(turkey Turkey) TurkeyAdapter {
	return TurkeyAdapter{
		turkey: turkey,
	}
}

// 这里需要实现鸭子的接口

func (adapter TurkeyAdapter) Quack() {
	adapter.turkey.Gobble()
}

func (adapter TurkeyAdapter) Fly() {
	for i := 0; i < 5; i++ {
		adapter.turkey.Fly()
	}
}
