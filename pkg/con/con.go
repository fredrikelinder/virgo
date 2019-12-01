package con

import "sync"

// For concurrently evaluates the given fn with values
// from 0 to n, and returns when all fn calls have completed.
func For(n int, fn func(i int)) {
	Spawn(n, fn).Wait()
}

// Spawn concurrently evaluates the given fn with values
// from 0 to n, and returns a WaitGroup to know when all
// fn calls have completed.
func Spawn(n int, fn func(i int)) *sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func(i int) {
			defer wg.Done()
			fn(i)
		}(i)
	}

	return &wg
}
