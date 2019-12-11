package ints

func Primes(n uint64) map[uint64]bool {
	switch {
	case n <= 2:
		return 0
	}
	prime := map[uint64]bool{2: true}
	for i := 3; i < n; i += 2 {
		prime[i] = true
	}
	var i uint64
	for i = 3; i < n; i += 2 {
		if prime[i] {
			for j := i * 2; j < n; j += i {
				delete(prime, j)
			}
		}
	}
	return prime
}
