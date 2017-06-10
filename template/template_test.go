package template

import (
	"testing"
)

func TestTemplate(t *testing.T) {
	var drink Preparer

	drink = NewCoffeeWithHook()
	drink.PrepareRecipe()
}