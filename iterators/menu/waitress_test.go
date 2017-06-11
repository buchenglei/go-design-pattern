package menu

import (
	"testing"
)

func TestWaitress(t *testing.T) {
	waitress := NewWaitress(
		NewPancakeHouseMenu(10),
		// 这里可以一次性添加多个菜单
	)

	waitress.PrintMenu()
}
