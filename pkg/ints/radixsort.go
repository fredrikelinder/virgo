package ints

import "math/bits"

// SortUint64s sorts the given uint64 slice in O(n) time.
func SortUint64s(nums []uint64) {
	const bits = 8
	const size = uint64(1) << bits

	tmps := make([]uint64, len(nums))
	max := Uint64sMax(nums...)

	for bit := 0; bit < 64 && max>>bit > 0; bit += bits {
		buckets := [size]uint{}

		// count number of respective buckets
		for _, n := range nums {
			buckets[(n>>bit)%size]++
		}

		// counts => positions
		for i := uint64(1); i < size; i++ {
			buckets[i] += buckets[i-1]
		}

		// copy value into tmps array
		for i := range nums {
			n := nums[len(nums)-i-1]
			buckets[(n>>bit)%size]--
			tmps[buckets[(n>>bit)%size]] = n
		}

		// copy value back into nums array
		copy(nums, tmps)
	}
}

// SortInt64s sorts the given int64 slice in O(n) time.
func SortInt64s(nums []int64) {
	const maxbit = uint64(1) << 63

	unums := int64sToUint64s(nums)

	for i, n := range nums {
		unums[i] = uint64(n) ^ maxbit
	}

	SortUint64s(unums)

	for i, n := range unums {
		nums[i] = int64(n ^ maxbit)
	}
}

// SortUint32s sorts the given uint32 slice in O(n) time.
func SortUint32s(nums []uint32) {
	const bits = 8
	const size = uint32(1) << bits

	tmps := make([]uint32, len(nums))
	max := Uint32sMax(nums...)

	for bit := 0; bit < 32 && max>>bit > 0; bit += bits {
		buckets := [size]uint{}

		// count number of respective buckets
		for _, n := range nums {
			buckets[(n>>bit)%size]++
		}

		// counts => positions
		for i := uint32(1); i < size; i++ {
			buckets[i] += buckets[i-1]
		}

		// copy value into tmps array
		for i := range nums {
			n := nums[len(nums)-i-1]
			buckets[(n>>bit)%size]--
			tmps[buckets[(n>>bit)%size]] = n
		}

		// copy value back into nums array
		copy(nums, tmps)
	}
}

// SortInt32s sorts the given int32 slice in O(n) time.
func SortInt32s(nums []int32) {
	const maxbit = uint32(1) << 31

	unums := int32sToUint32s(nums)

	for i, n := range nums {
		unums[i] = uint32(n) ^ maxbit
	}

	SortUint32s(unums)

	for i, n := range unums {
		nums[i] = int32(n ^ maxbit)
	}
}

// SortUint16s sorts the given uint16 slice in O(n) time.
func SortUint16s(nums []uint16) {
	const bits = 8
	const size = uint16(1) << bits

	tmps := make([]uint16, len(nums))
	max := Uint16sMax(nums...)

	for bit := 0; bit < 16 && max>>bit > 0; bit += bits {
		buckets := [size]uint{}

		// count number of respective buckets
		for _, n := range nums {
			buckets[(n>>bit)%size]++
		}

		// counts => positions
		for i := uint16(1); i < size; i++ {
			buckets[i] += buckets[i-1]
		}

		// copy value into tmps array
		for i := range nums {
			n := nums[len(nums)-i-1]
			buckets[(n>>bit)%size]--
			tmps[buckets[(n>>bit)%size]] = n
		}

		// copy value back into nums array
		copy(nums, tmps)
	}
}

// SortInt16s sorts the given int16 slice in O(n) time.
func SortInt16s(nums []int16) {
	const maxbit = uint16(1) << 15

	unums := int16sToUint16s(nums)

	for i, n := range nums {
		unums[i] = uint16(n) ^ maxbit
	}

	SortUint16s(unums)

	for i, n := range unums {
		nums[i] = int16(n ^ maxbit)
	}
}

// SortUint8s sorts the given uint8 slice in O(n) time.
func SortUint8s(nums []uint8) {
	const bits = 4
	const size = uint8(1) << bits

	tmps := make([]uint8, len(nums))
	max := Uint8sMax(nums...)

	for bit := 0; bit < 8 && max>>bit > 0; bit += bits {
		buckets := [size]uint{}

		// count number of respective buckets
		for _, n := range nums {
			buckets[(n>>bit)%size]++
		}

		// counts => positions
		for i := uint8(1); i < size; i++ {
			buckets[i] += buckets[i-1]
		}

		// copy value into tmps array
		for i := range nums {
			n := nums[len(nums)-i-1]
			buckets[(n>>bit)%size]--
			tmps[buckets[(n>>bit)%size]] = n
		}

		// copy value back into nums array
		copy(nums, tmps)
	}
}

// SortInt8s sorts the given int8 slice in O(n) time.
func SortInt8s(nums []int8) {
	const maxbit = uint8(1) << 7

	unums := int8sToUint8s(nums)

	for i, n := range nums {
		unums[i] = uint8(n) ^ maxbit
	}

	SortUint8s(unums)

	for i, n := range unums {
		nums[i] = int8(n ^ maxbit)
	}
}

// SortUints sorts the given uint slice in O(n) time.
func SortUints(nums []uint) {
	switch bits.UintSize {
	case 32:
		unums := uintsToUint32s(nums)
		SortUint32s(unums)
		uint32sToUints(nums, unums)
	case 64:
		unums := uintsToUint64s(nums)
		SortUint64s(unums)
		uint64sToUints(nums, unums)
	default:
		panic("cannot convert from []uint")
	}
}

// SortInts sorts the given int slice in O(n) time.
func SortInts(nums []int) {
	switch bits.UintSize {
	case 32:
		unums := intsToInt32s(nums)
		SortInt32s(unums)
		int32sToInts(nums, unums)
	case 64:
		unums := intsToInt64s(nums)
		SortInt64s(unums)
		int64sToInts(nums, unums)
	default:
		panic("cannot convert from []int")
	}
}
