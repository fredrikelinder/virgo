package ints

// IsSortedUint64s returns true if the given nums are sorted,
// and false otherwise.
func IsSortedUint64s(nums []uint64) bool {
	if len(nums) == 0 {
		return true
	}

	prev := nums[0]

	for _, n := range nums[1:] {
		if prev > n {
			return false
		}
		prev = n
	}

	return true
}

// IsSortedInt64s returns true if the given nums are sorted,
// and false otherwise.
func IsSortedInt64s(nums []int64) bool {
	if len(nums) == 0 {
		return true
	}

	prev := nums[0]

	for _, n := range nums[1:] {
		if prev > n {
			return false
		}
		prev = n
	}

	return true
}

// IsSortedUint32s returns true if the given nums are sorted,
// and false otherwise.
func IsSortedUint32s(nums []uint32) bool {
	if len(nums) == 0 {
		return true
	}

	prev := nums[0]

	for _, n := range nums[1:] {
		if prev > n {
			return false
		}
		prev = n
	}

	return true
}

// IsSortedInt32s returns true if the given nums are sorted,
// and false otherwise.
func IsSortedInt32s(nums []int32) bool {
	if len(nums) == 0 {
		return true
	}

	prev := nums[0]

	for _, n := range nums[1:] {
		if prev > n {
			return false
		}
		prev = n
	}

	return true
}

// IsSortedUint16s returns true if the given nums are sorted,
// and false otherwise.
func IsSortedUint16s(nums []uint16) bool {
	if len(nums) == 0 {
		return true
	}

	prev := nums[0]

	for _, n := range nums[1:] {
		if prev > n {
			return false
		}
		prev = n
	}

	return true
}

// IsSortedInt16s returns true if the given nums are sorted,
// and false otherwise.
func IsSortedInt16s(nums []int16) bool {
	if len(nums) == 0 {
		return true
	}

	prev := nums[0]

	for _, n := range nums[1:] {
		if prev > n {
			return false
		}
		prev = n
	}

	return true
}

// IsSortedUint8s returns true if the given nums are sorted,
// and false otherwise.
func IsSortedUint8s(nums []uint8) bool {
	if len(nums) == 0 {
		return true
	}

	prev := nums[0]

	for _, n := range nums[1:] {
		if prev > n {
			return false
		}
		prev = n
	}

	return true
}

// IsSortedInt8s returns true if the given nums are sorted,
// and false otherwise.
func IsSortedInt8s(nums []int8) bool {
	if len(nums) == 0 {
		return true
	}

	prev := nums[0]

	for _, n := range nums[1:] {
		if prev > n {
			return false
		}
		prev = n
	}

	return true
}

// IsSortedUints returns true if the given nums are sorted,
// and false otherwise.
func IsSortedUints(nums []uint) bool {
	if len(nums) == 0 {
		return true
	}

	prev := nums[0]

	for _, n := range nums[1:] {
		if prev > n {
			return false
		}
		prev = n
	}

	return true
}

// IsSortedInts returns true if the given nums are sorted,
// and false otherwise.
func IsSortedInts(nums []int) bool {
	if len(nums) == 0 {
		return true
	}

	prev := nums[0]

	for _, n := range nums[1:] {
		if prev > n {
			return false
		}
		prev = n
	}

	return true
}
