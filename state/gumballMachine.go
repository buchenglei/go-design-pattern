package state

import "fmt"

// GumbalMachine 定义一个糖果机
type GumbalMachine struct {
	SoldOutState    State
	NoQuarterState  State
	HasQuarterState State
	SoldState       State

	state State
	count int
}

func NewGumbalMachine(numberGumballs int) *GumbalMachine {
	gm := &GumbalMachine{
		count: numberGumballs,
	}

	gm.NoQuarterState = &NoQuarterState{gm: gm}
	gm.HasQuarterState = &HasQuarterState{gm: gm}
	gm.SoldState = &SoldState{gm: gm}
	gm.SoldOutState = &SoldOutState{gm: gm}

	if numberGumballs > 0 {
		gm.state = gm.NoQuarterState
	} else {
		gm.state = gm.SoldOutState
	}

	return gm
}

func (gm *GumbalMachine) InsertQuarter() {
	gm.state.InsertQuarter()
}

func (gm *GumbalMachine) EjectQuarter() {
	gm.state.EjectQuartar()
}

func (gm *GumbalMachine) TurmCrank() {
	gm.state.TurnCrank()
	gm.state.Dispense()
}

func (gm *GumbalMachine) SetState(state State) {
	gm.state = state
}

func (gm *GumbalMachine) ReleaseBall() {
	fmt.Println("A gumball comes rolling out the slot...")
	if gm.count != 0 {
		gm.count--
	}
}

func (gm *GumbalMachine) GetCount() int {
	return gm.count
}
