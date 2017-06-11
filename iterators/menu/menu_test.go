package menu

import (
	"fmt"
	"testing"
)

func TestMenu(t *testing.T) {
	var menu Menu
	menu = NewPancakeHouseMenu(10)

	iter := menu.CreateItearator()
	for iter.HasNext() {
		fmt.Printf("MenuItem: %+v\n", iter.Next())
	}
}
