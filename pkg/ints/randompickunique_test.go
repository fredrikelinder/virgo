package ints

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestRandPickUnique(t *testing.T) {
	cases := []struct {
		size int
		n    int
	}{
		{size: 10, n: 2},
		{size: 100, n: 20},
		{size: 1000, n: 200},
	}

	for i, c := range cases {
		for _, strategy := range strategies() {
			nums := genRandNums(c.size, strategy)
			name := fmt.Sprintf("%v-%v", i, strategy)
			t.Run(name, func(t *testing.T) {
				picked, ok := RandPickUnique(nums, c.n)
				if !ok {
					return
				}
				m := map[int]int{}
				for _, n := range nums {
					m[n] += 1
				}
				for _, p := range picked {
					if _, ok := m[p]; !ok {
						t.Errorf("%v: %+v vs %+v", i, picked, m)
					}
				}
			})
		}
	}
}

func BenchmarkRandPickUnique(b *testing.B) {
	cases := []struct {
		size int
		n    int
	}{
		{size: 10, n: 2},
		{size: 100, n: 20},
		{size: 1000, n: 200},
	}

	for i, c := range cases {
		for _, strategy := range strategies() {
			nums := genRandNums(c.size, strategy)
			name := fmt.Sprintf("%v-%v", i, strategy)
			b.Run(name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					RandPickUnique(nums, c.n)
				}
			})
		}
	}
}

func strategies() []strategy {
	return []strategy{Sparse, Average, Dense, Superdense}
}

type strategy int

const (
	Sparse strategy = iota
	Average
	Dense
	Superdense
)

func genRandNums(n int, s strategy) []int {
	nums := make([]int, n)

	var k float64
	switch s {
	case Sparse:
		k = 0.1
	case Average:
		k = 0.5
	case Dense:
		k = 0.9
	case Superdense:
		nums[len(nums)-1] = 1
		return nums
	}

	for i := 0; i < n; i++ {
		r := rand.Uint64()
		if float64(r)/float64(MaxInt64) < k {
			nums[i] = int(rand.Uint64())
		}
	}

	return nums
}
