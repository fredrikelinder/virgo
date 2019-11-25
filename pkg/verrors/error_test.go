package verrors

import (
	"errors"
	"fmt"
)

func ExampleError_stack() {
	second := func() error {
		err := errors.New("third - the original error")
		return Errorf(err, "[%@] second\n")
	}

	first := func() error {
		return Errorf(second(), "[%@] first\n")
	}

	fmt.Printf("%v\n", first())
	// Output:
	// [verrors/error_test.go:15] first
	// [verrors/error_test.go:11] second
	// third - the original error
}

func ExampleError_nil() {
	fmt.Printf("%v\n", Errorf(nil, "[%@] nil error"))
	// Output:
	// <nil>
}
