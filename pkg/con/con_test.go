package con

import "fmt"

func Example_For() {
	values := make([]int, 4)

	For(len(values), func(i int) {
		values[i] = i
	})

	fmt.Println(values)

	// Output:
	// [0 1 2 3]
}
