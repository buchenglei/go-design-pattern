package state

import "testing"

func TestGunmballMachine(t *testing.T) {
	gm := NewGumbalMachine(5)

	gm.InsertQuarter()
	gm.TurmCrank()

}
