package menu

const maxtMenuItemCount = 20

// Menu 定义菜单需要实现的方法
type Menu interface {
	// 为了简化示例，这里菜单只需要获取迭代器即可
	// 实际场景中看实际需要
	CreateItearator() Interator
}

type PancakeHouseMenu struct {
	items []*MenuItem
}

func NewPancakeHouseMenu(itemCount int) *PancakeHouseMenu {
	menu := &PancakeHouseMenu{}

	if itemCount <= 0 {
		itemCount = maxtMenuItemCount
	}
	menu.items = make([]*MenuItem, itemCount)

	// 添加菜单项
	menu.items[0] = NewMenuItem(
		"K&B's Pancake Breakfast",
		"Pancakes with scrambled eggs, and toast",
		true,
		2.99,
	)

	menu.items[1] = NewMenuItem(
		"Regular Pancake Breakfast",
		"Pancakes with fried eggs, sausage",
		false,
		2.99,
	)

	menu.items[2] = NewMenuItem(
		"Blueberry Pancakes",
		"Pancakes mode with fresh blueberries",
		true,
		3.49,
	)

	menu.items[3] = NewMenuItem(
		"Waffles",
		"Waffles, with your choice of blueberries or trawberries",
		true,
		3.59,
	)

	return menu
}

// CreateItearator 返回一个迭代器
func (menu *PancakeHouseMenu) CreateItearator() Interator {
	return NewPancakeHouseMenuIter(menu)
}
