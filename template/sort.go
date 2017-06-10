package template

import "fmt"
import "sort"

// Duck 定义一个基本的鸭子类型
type Duck struct {
	name   string
	weight int
}

// SomeDucks 定义一组鸭子类型
type SomeDucks []Duck

/***** 一组鸭子类型需要实现排序接口的方法 ******/

// go标准库中的排序接口
// type Interface interface {
//     // Len is the number of elements in the collection.
//     Len() int
//     // Less reports whether the element with
//     // index i should sort before the element with index j.
//     Less(i, j int) bool
//     // Swap swaps the elements with indexes i and j.
//     Swap(i, j int)
// }

// Len 实现Len方法
func (ducks SomeDucks) Len() int {
	return len(ducks)
}

// Less 实现Less方法
func (ducks SomeDucks) Less(i, j int) bool {
	if ducks[i].weight <= ducks[j].weight {
		return true
	}

	return false
}

func (ducks SomeDucks) Swap(i, j int) {
	ducks[j], ducks[i] = ducks[i], ducks[j]
}

// SortDucks 对鸭子进行排序
func SortDucks(ducks SomeDucks) {
	fmt.Println("Is sorted: ", sort.IsSorted(ducks))
	fmt.Println("Will sort.....")

	// 在这个方法中并不需要知道标准库的排序算法是什么样的
	// 通过接口定义的方法，实现了从外部干预内部算法的目的
	sort.Sort(ducks)
	fmt.Printf("After sort: %+v\n", ducks)
	fmt.Println("Is sorted: ", sort.IsSorted(ducks))
}
