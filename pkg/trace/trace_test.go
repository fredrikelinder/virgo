package trace

import "fmt"

func f() {
	defer Exit(Enter(fmt.Printf))
}

func Example_Enter() {
	f()

	// Output:
	// > trace/trace_test.go:6 f
	// < trace/trace_test.go:6 f
}
