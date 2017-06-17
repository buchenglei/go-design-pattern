package state

import "fmt"

type State interface {
	InsertQuarter()
	EjectQuartar()
	TurnCrank()
	Dispense()
}

/****** 实现不同的状态 *****/

type NoQuarterState struct {
	gm *GumbalMachine
}

func (noquater *NoQuarterState) InsertQuarter() {
	fmt.Println("Current state is NoQuarterState, You inserted a quarter")
	noquater.gm.SetState(noquater.gm.HasQuarterState)
}

func (NoQuarterState) EjectQuartar() {
	fmt.Println("Current state is NoQuarterState,You haven't inserted a quarter")
}

func (NoQuarterState) TurnCrank() {
	fmt.Println("Current state is NoQuarterState,You turned, but there's no quarter")
}

func (NoQuarterState) Dispense() {
	fmt.Println("Current state is NoQuarterState,You need to pay first")
}

type HasQuarterState struct {
	gm *GumbalMachine
}

func (HasQuarterState) InsertQuarter() {
	fmt.Println("Current state is HasQuarterState,You can't inserter another quarter")
}

func (hasQuarter HasQuarterState) EjectQuartar() {
	fmt.Println("Current state is HasQuarterState, Quarter returned")
	hasQuarter.gm.SetState(hasQuarter.gm.NoQuarterState)
}

func (hasQuarter *HasQuarterState) TurnCrank() {
	fmt.Println("Current state is HasQuarterState,You turned.....")
	hasQuarter.gm.SetState(hasQuarter.gm.SoldState)
}

func (HasQuarterState) Dispense() {
	fmt.Println("Current state is HasQuarterState,No gumball dispensed")
}

type SoldState struct {
	gm *GumbalMachine
}

func (SoldState) InsertQuarter() {
	fmt.Println("Current state is SoldState,Please wait, we're already giving you a gumball")
}

func (SoldState) EjectQuartar() {
	fmt.Println("Current state is SoldState,Sorry, you already turned the crank")
}

func (SoldState) TurnCrank() {
	fmt.Println("Current state is SoldState,Turning twice doesn't get you another gumball!")
}

func (sold SoldState) Dispense() {
	sold.gm.ReleaseBall()
	if sold.gm.GetCount() > 0 {
		sold.gm.SetState(sold.gm.NoQuarterState)
	} else {
		fmt.Println("Current state is SoldState,Oops, out of gumballs")
		sold.gm.SetState(sold.gm.SoldOutState)
	}
}

type SoldOutState struct {
	gm *GumbalMachine
}

func (SoldOutState) InsertQuarter() {
	fmt.Println("Current state is SoldOutState,Sorry, there is no gumball")
}

func (SoldOutState) EjectQuartar() {
	fmt.Println("Current state is SoldOutState,Sorry, there is no gumball")
}

func (SoldOutState) TurnCrank() {
	fmt.Println("Current state is SoldOutState,Sorry, there is no gumball")
}

func (SoldOutState) Dispense() {
	fmt.Println("Current state is SoldOutState,Sorry, there is no gumball")
}
