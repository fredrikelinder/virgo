package fsm

import (
	"errors"

	"github.com/fredrikelinder/virgo/pkg/verrors"
)

var (
	// ErrCannotReinitialize is returned when calling Init a second time.
	ErrCannotReinitialize = errors.New("cannot re-initialize")

	// ErrMustInitializeFirst is returned when Init is not the first method being called.
	ErrMustInitializeFirst = errors.New("must initialize first")

	// ErrTransitionsCannotBeNil is returned when trying to Init without transition.
	ErrTransitionsCannotBeNil = errors.New("transitions cannot be nil")

	// ErrUnacceptableInitialState is returned when trying to Init with an initial state
	// that is not present in the transition.
	ErrUnacceptableInitialState = errors.New("unacceptable initial state")

	// ErrUnacceptableEvent is returned when calling On with an Event
	// that has no transition from the current state.
	ErrUnacceptableEvent = errors.New("unacceptable event ID")

	// ErrUnacceptableStateCb is returned when calling Init with a state callback
	// for a state that does not exist in the transitions.
	ErrUnacceptableStateCb = errors.New("unacceptable state callback")

	// ErrUnacceptableStateTransition is returned when calling the To action
	// trying to transition to a state that is not present in the transitions.
	ErrUnacceptableStateTransition = errors.New("unacceptable state transition")
)

// EventID uniquely identifies an event.
//
// Tip: use stringer to convert your constants to string.
type EventID int

// Event represents incoming events.
type Event interface {
	ID() EventID
}

// Action performs some action using the given event.
type Action func(event Event) error

// StateID represents a state in the FSM.
type StateID int

// Transitions defines the possible transitions in an FSM.
type Transitions map[StateID]map[EventID][]Action

// StateCb defines callbacks when entering and leaving a state.
// The Enter and leave actions will be called with a nil event.
type StateCb struct {
	Enter Action
	Leave Action
}

// StateCbs defines the possible state callbacks in an FSM.
type StateCbs map[StateID]StateCb

// FSM represents a finite state machine.
type FSM struct {
	hot                   []Event
	postponed             []Event
	postponedTransitionID int
	state                 StateID
	stateCbs              StateCbs
	transitions           Transitions
}

type noError struct{}

func (*noError) Error() string { return "" }

type postponeNoError struct {
	noError
}

type stateTransitionNoError struct {
	noError
	to StateID
}

// New returns a new FSM.
func New() *FSM {
	return &FSM{}
}

// Init initializes the FSM, to start in the initial state, having
// the given transitions, and state callbacks.
//
// Note: transitions cannot be nil.
// Note: stateCbs can be nil.
func (f *FSM) Init(initialState StateID, transitions Transitions, stateCbs StateCbs) error {
	if f.transitions != nil {
		return ErrCannotReinitialize
	}

	if transitions == nil {
		return ErrTransitionsCannotBeNil
	}

	if _, ok := transitions[initialState]; !ok {
		return verrors.Errorf(ErrUnacceptableInitialState, "state: %v", initialState)
	}

	if stateCbs == nil {
		stateCbs = map[StateID]StateCb{}
	}

	for stateID := range stateCbs {
		if _, ok := transitions[stateID]; !ok {
			return verrors.Errorf(ErrUnacceptableStateCb, "state: %v", stateID)
		}
	}

	f.state = initialState
	f.stateCbs = stateCbs
	f.transitions = transitions

	// We're entering the new state, call it's enter callback
	cbs, ok := f.stateCbs[f.state]
	if ok && cbs.Enter != nil {
		err := cbs.Enter(nil)
		if err != nil {
			return verrors.Errorf(err, "enter callback failed")
		}
	}

	return nil
}

// On applies the given event on the FSM.
//
// If an error is returned, then the FSM may be in an unusable state.
func (f *FSM) On(event Event) error {
	if f.transitions == nil {
		return ErrMustInitializeFirst
	}

	actions, ok := f.transitions[f.state][event.ID()]
	if !ok {
		return verrors.Errorf(ErrUnacceptableEvent, "no such event: %v", event.ID())
	}

	for i, action := range actions {
		err := action(event)
		if err != nil {
			switch x := err.(type) {
			case *postponeNoError:
				return nil
			case *stateTransitionNoError:
				err := f.stateTransitionLeave()
				if err != nil {
					return err
				}

				f.state = x.to

				err = f.stateTransitionEnter()
				if err != nil {
					return err
				}

				err = f.stateTransitionPostponed()
				if err != nil {
					return err
				}

				return nil
			default:
				return verrors.Errorf(err, "action %v failed", i)
			}
		}
	}

	return nil
}

func (f *FSM) stateTransitionLeave() error {
	cbs, ok := f.stateCbs[f.state]
	if ok && cbs.Leave != nil {
		err := cbs.Leave(nil)
		if err != nil {
			return verrors.Errorf(err, "leave callback failed")
		}
	}

	return nil
}

func (f *FSM) stateTransitionEnter() error {
	cbs, ok := f.stateCbs[f.state]
	if ok && cbs.Enter != nil {
		err := cbs.Enter(nil)
		if err != nil {
			return verrors.Errorf(err, "enter callback failed")
		}
	}

	return nil
}

func (f *FSM) stateTransitionPostponed() error {
	f.postponedTransitionID++
	transitionID := f.postponedTransitionID

	f.hot = append(f.hot, f.postponed...)
	f.postponed = nil

	for transitionID == f.postponedTransitionID && len(f.hot) > 0 {
		var event Event
		event, f.hot = f.hot[0], f.hot[1:]

		err := f.On(event)
		if err != nil {
			return verrors.Errorf(err, "postponed action failed")
		}
	}

	return nil
}

// Postpone pauses the evaluation of the event until
// the next state transition. Any actions after this
// action are ignored.
func (f *FSM) Postpone(event Event) error {
	f.postponed = append(f.postponed, event)
	return &postponeNoError{}
}

// PostponeTo evaluates the event after transitioning
// to the given state. Any actions after this action
// are ignored.
func (f *FSM) PostponeTo(state StateID) Action {
	return func(event Event) error {
		_, ok := f.transitions[state]
		if !ok {
			return verrors.Errorf(ErrUnacceptableEvent, "cannot transition to non-existing state: %v", state)
		}

		f.postponed = append(f.postponed, event)
		return &stateTransitionNoError{to: state}
	}
}

// To transitions to the given state. Any actions after
// this action are ignored.
func (f *FSM) To(state StateID) Action {
	return func(Event) error {
		_, ok := f.transitions[state]
		if !ok {
			verrors.Errorf(ErrUnacceptableStateTransition, "cannot transition to non-existing state: %v", state)
		}

		return &stateTransitionNoError{to: state}
	}
}
