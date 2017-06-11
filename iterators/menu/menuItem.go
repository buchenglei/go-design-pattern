package menu

// MenuItem 定义基本的菜单项
type MenuItem struct {
	name       string
	desc       string
	vegetarian bool
	price      float32
}

func NewMenuItem(name, desc string, vegetarian bool, price float32) *MenuItem {
	return &MenuItem{
		name:       name,
		desc:       desc,
		vegetarian: vegetarian,
		price:      price,
	}
}

func (item MenuItem) GetName() string {
	return item.name
}

func (item MenuItem) GetDesc() string {
	return item.desc
}

func (item MenuItem) Price() float32 {
	return item.price
}

func (item MenuItem) IsVegetarian() bool {
	return item.vegetarian
}
