package verrors

import (
	"errors"
	"fmt"
)

func ExampleErrors_stack() {
	one := errors.New("one")
	two := errors.New("two")
	fmt.Printf("%v\n", Errorsf([]error{one, two}, "a bunch of errors:\n"))

	// Output:
	// a bunch of errors:
	// [one, two]
}

func ExampleErrors_nil() {
	fmt.Printf("%v\n", Errorsf(nil, "a bunch of errors:\n"))

	// Output:
	// <nil>
}
