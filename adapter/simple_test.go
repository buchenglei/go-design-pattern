package adapter

import (
	"testing"
)

func TestTurkeyAdapter(t *testing.T) {
	testDuck := func(duck Duck) {
		duck.Quack()
		duck.Fly()
	}

	duck := MallarDuck{}
	turkey := WildTurkey{}

	t.Log("The WildTurkey says: ")
	turkey.Gobble()
	turkey.Fly()

	testDuck(duck)

	// 这实际上是一只鸭子
	turkeyAdapter := NewTurkeyAdapter(turkey)

	testDuck(turkeyAdapter)
}
