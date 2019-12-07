package con

import "fmt"

func ExampleCurrent() {
	values := make([]int, 4)

	Current(len(values), func(i int) {
		values[i] = i
	})

	fmt.Println(values)

	// Output:
	// [0 1 2 3]
}

func ExampleCurrent_limitedConcurrency() {
	values := make([]int, 20)

	// sequence using a channel
	ch := make(chan int)

	// use a separate goroutine to send
	go func() {
		defer close(ch)
		for i := range values {
			ch <- i
		}
	}()

	// use Current to limit number of concurrent receivers
	Current(2, func(i int) {
		for v := range ch {
			values[i] = i
		}
	})

	fmt.Println(values)
	// Output:
	// [0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19]
}
