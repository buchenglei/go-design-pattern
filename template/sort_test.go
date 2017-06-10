package template

import "testing"

func TestSortDucks(t *testing.T) {
	ducks := SomeDucks{
		Duck{name: "Z", weight: 100},
		Duck{name: "F", weight: 78},
		Duck{name: "B", weight: 89},
		Duck{name: "A", weight: 21},
		Duck{name: "W", weight: 77},
	}

	SortDucks(ducks)
}
