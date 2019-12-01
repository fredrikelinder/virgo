// +build !unsafe

package ints

func int64sToUint64s(nums []int64) []uint64 {
	const maxbit = uint64(1) << 63

	unums := make([]uint64, len(nums))

	for i, n := range nums {
		unums[i] = uint64(n) ^ maxbit
	}

	return unums
}

func int32sToUint32s(nums []int32) []uint32 {
	const maxbit = uint32(1) << 31

	unums := make([]uint32, len(nums))

	for i, n := range nums {
		unums[i] = uint32(n) ^ maxbit
	}

	return unums
}

func int16sToUint16s(nums []int16) []uint16 {
	const maxbit = uint16(1) << 15

	unums := make([]uint16, len(nums))

	for i, n := range nums {
		unums[i] = uint16(n) ^ maxbit
	}

	return unums
}

func int8sToUint8s(nums []int8) []uint8 {
	const maxbit = uint8(1) << 7

	unums := make([]uint8, len(nums))

	for i, n := range nums {
		unums[i] = uint8(n) ^ maxbit
	}

	return unums
}

func uintsToUint64s(nums []uint) []uint64 {
	unums := make([]uint64, len(nums))

	for i, n := range nums {
		unums[i] = uint64(n)
	}

	return unums
}

func intsToInt64s(nums []int) []int64 {
	unums := make([]int64, len(nums))

	for i, n := range nums {
		unums[i] = int64(n)
	}

	return unums
}

func uintsToUint32s(nums []uint) []uint32 {
	unums := make([]uint32, len(nums))

	for i, n := range nums {
		unums[i] = uint32(n)
	}

	return unums
}

func intsToInt32s(nums []int) []int32 {
	unums := make([]int32, len(nums))

	for i, n := range nums {
		unums[i] = int32(n)
	}

	return unums
}

func uint64sToUints(nums []uint, unums []uint64) {
	for i, n := range unums {
		nums[i] = uint(n)
	}
}

func int64sToInts(nums []int, unums []int64) {
	for i, n := range unums {
		nums[i] = int(n)
	}
}

func uint32sToUints(nums []uint, unums []uint32) {
	for i, n := range unums {
		nums[i] = uint(n)
	}
}

func int32sToInts(nums []int, unums []int32) {
	for i, n := range unums {
		nums[i] = int(n)
	}
}
