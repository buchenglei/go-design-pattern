package menu

import "fmt"

type Waitress struct {
	menus []Menu
}

func NewWaitress(menus ...Menu) *Waitress {
	return &Waitress{
		menus: menus,
	}
}

func (waitress *Waitress) AddMenu(menu Menu) {
	waitress.menus = append(waitress.menus, menu)
}

func (waitress Waitress) PrintMenu() {
	var iter Interator

	for i, menu := range waitress.menus {
		iter = menu.CreateItearator()
		fmt.Printf("第%d份菜单\nMenuItem:\n", i+1)

		for iter.HasNext() {
			item := iter.Next()
			fmt.Printf(
				"Name: %s\n\tvegetarian:%v\n\tprice:%.2f\n\tdesc:%s\n",
				item.GetName(),
				item.IsVegetarian(),
				item.Price(),
				item.GetDesc(),
			)
		}
	}
}
