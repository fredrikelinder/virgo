package fsm

import (
	"fmt"
)

//go:generate stringer -type ExampleStateID
type ExampleStateID int

//go:generate stringer -type ExampleEventID
type ExampleEventID int

const (
	stateZero = /*ExampleStateID*/ iota
	stateOne
	stateTwo
	stateRed
	stateGreen
	stateYellow
)

const (
	eventTick = /*ExampleEventID*/ iota
	eventPostpone
	eventPostponeTo
	eventAlea
	eventIacta
	eventEst
)

type event struct {
	id EventID
}

func (e *event) ID() EventID { return e.id }

func makePrinter(lights string) func(Event) error {
	return func(Event) error {
		fmt.Println(lights)
		return nil
	}
}

func ExampleFSM_actions() {
	f := New()
	pie(f.Init(stateRed, Transitions{
		stateRed:    {eventTick: {makePrinter("⚫⚫🍎"), f.To(stateGreen)}},
		stateGreen:  {eventTick: {makePrinter("🍏⚫⚫"), f.To(stateYellow)}},
		stateYellow: {eventTick: {makePrinter("🍏🍋⚫"), f.To(stateRed)}},
	}, nil))

	tick := &event{id: eventTick}
	pie(f.On(tick))
	pie(f.On(tick))
	pie(f.On(tick))
	pie(f.On(tick))

	// Output:
	// ⚫⚫🍎
	// 🍏⚫⚫
	// 🍏🍋⚫
	// ⚫⚫🍎
}

func ExampleFSM_stateCbs() {
	f := New()
	pie(f.Init(stateRed, Transitions{
		stateRed:    {eventTick: {f.To(stateGreen)}},
		stateGreen:  {eventTick: {f.To(stateYellow)}},
		stateYellow: {eventTick: {f.To(stateRed)}},
	}, StateCbs{
		stateRed:    {Enter: makePrinter("⚫⚫🍎")},
		stateGreen:  {Enter: makePrinter("🍏⚫⚫")},
		stateYellow: {Enter: makePrinter("🍏🍋⚫")},
	}))

	tick := &event{id: eventTick}
	pie(f.On(tick))
	pie(f.On(tick))
	pie(f.On(tick))

	// Output:
	// ⚫⚫🍎
	// 🍏⚫⚫
	// 🍏🍋⚫
	// ⚫⚫🍎
}

func ExampleFSM_postpone() {
	ticker := makePrinter("tick")
	postponed := makePrinter("postponed")

	f := New()
	pie(f.Init(stateZero, Transitions{
		stateZero: {
			eventPostpone: {f.Postpone},
			eventTick:     {ticker, f.To(stateOne)},
		},
		stateOne: {
			eventPostpone:   {postponed},
			eventPostponeTo: {ticker, f.PostponeTo(stateTwo)},
		},
		stateTwo: {
			eventPostponeTo: {postponed},
		},
	}, nil))

	tick := &event{id: eventTick}
	postpone := &event{id: eventPostpone}
	postponeTo := &event{id: eventPostponeTo}
	pie(f.On(postpone))
	fmt.Println("-0-")
	pie(f.On(tick))
	fmt.Println("-1-")
	pie(f.On(postponeTo))

	// Output:
	// -0-
	// tick
	// postponed
	// -1-
	// tick
	// postponed
}

func ExampleFSM_postponeDeep() {
	f := New()
	pie(f.Init(stateZero, Transitions{
		stateZero: {
			eventAlea:  {makePrinter("alea"), f.To(stateOne)},
			eventIacta: {f.Postpone},
			eventEst:   {f.Postpone},
		},
		stateOne: {
			eventAlea:  {f.Postpone},
			eventIacta: {makePrinter("iacta"), f.To(stateTwo)},
			eventEst:   {f.Postpone},
		},
		stateTwo: {
			eventAlea:  {f.Postpone},
			eventIacta: {f.Postpone},
			eventEst:   {makePrinter("est"), f.To(stateZero)},
		},
	}, nil))

	alea := &event{id: eventAlea}
	iacta := &event{id: eventIacta}
	est := &event{id: eventEst}
	pie(f.On(est))
	fmt.Println("-0-")
	pie(f.On(est))
	fmt.Println("-1-")
	pie(f.On(iacta))
	fmt.Println("-2-")
	pie(f.On(est))
	fmt.Println("-3-")
	pie(f.On(alea))
	fmt.Println("-4-")
	pie(f.On(iacta))
	fmt.Println("-5-")
	pie(f.On(alea))

	// Output:
	// -0-
	// -1-
	// -2-
	// -3-
	// alea
	// iacta
	// est
	// -4-
	// -5-
	// alea
	// iacta
	// est
}

func pie(err error) {
	if err != nil {
		fmt.Printf("err: %#v\n", err)
	}
}
