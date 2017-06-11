package menu

// Interator 定义菜单项的迭代器
type Interator interface {
	HasNext() bool
	Next() *MenuItem
}

type PancakeHouseMenuIter struct {
	menu *PancakeHouseMenu
	pos  int
}

// NewPancakeHouseMenuIter 创建一个PancakeHouseMenu的迭代器
func NewPancakeHouseMenuIter(menu *PancakeHouseMenu) *PancakeHouseMenuIter {
	return &PancakeHouseMenuIter{
		menu: menu,
		pos:  0,
	}
}

// HasNext 判断PancakeHouseMenu是否还有未完的迭代项目
func (iter PancakeHouseMenuIter) HasNext() bool {
	// PancakeHouseMenu 使用的是预分配的slice
	if iter.pos >= len(iter.menu.items)-1 {
		// 避免越界
		return false
	}

	if iter.menu.items[iter.pos+1] == nil {
		return false
	}

	return true
}

// Next 返回PancakeHouseMenu的下一个元素
func (iter *PancakeHouseMenuIter) Next() *MenuItem {
	iter.pos++
	return iter.menu.items[iter.pos]
}
