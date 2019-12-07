package ints

import "math/rand"

// RandPickUnique returns the given number of unique numbers,
// randomly picked from the given list. It will return true
// if the given list has enough unique numbers.
func RandPickUnique(nums []int, n int) ([]int, bool) {
	m := map[int]bool{}

	for len(nums) > 0 && len(m) < n {
		i := rand.Uint32() % uint32(len(nums))

		m[nums[i]] = true

		last := len(nums) - 1
		nums[i], nums[last] = nums[last], nums[i]
		nums = nums[:last]
	}

	var vals []int
	for v := range m {
		vals = append(vals, v)
	}

	return vals, len(vals) == n
}
