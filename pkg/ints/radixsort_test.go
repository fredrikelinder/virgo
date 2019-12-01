package ints

import (
	"math/rand"
	"strconv"
	"testing"
)

func TestSortUint64s(t *testing.T) {
	for i := 0; i < 4; i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			nums := generateUint64s(100)
			SortUint64s(nums)
			if !IsSortedUint64s(nums) {
				t.Errorf("unsorted %b", nums)
			}
		})
	}
}

func TestSortInt64s(t *testing.T) {
	for i := 0; i < 4; i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			nums := generateInt64s(100)
			SortInt64s(nums)
			if !IsSortedInt64s(nums) {
				t.Errorf("unsorted %b", nums)
			}
		})
	}
}

func BenchmarkSortUint64s(b *testing.B) {
	for i := 0; i < 1; i++ {
		b.Run(strconv.Itoa(i), func(b *testing.B) {
			nums := generateUint64s(b.N)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				SortUint64s(nums)
			}
		})
	}
}

func BenchmarkSortInt64s(b *testing.B) {
	for i := 0; i < 1; i++ {
		b.Run(strconv.Itoa(i), func(b *testing.B) {
			nums := generateInt64s(b.N)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				SortInt64s(nums)
			}
		})
	}
}

func generateUint64s(n int) []uint64 {
	nums := make([]uint64, n)
	for i := range nums {
		nums[i] = rand.Uint64()
	}
	return nums
}

func generateInt64s(n int) []int64 {
	nums := make([]int64, n)
	for i := range nums {
		nums[i] = int64(rand.Uint64())
	}
	return nums
}

func TestSortUint32s(t *testing.T) {
	for i := 0; i < 4; i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			nums := generateUint32s(100)
			SortUint32s(nums)
			if !IsSortedUint32s(nums) {
				t.Errorf("unsorted %b", nums)
			}
		})
	}
}

func TestSortInt32s(t *testing.T) {
	for i := 0; i < 4; i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			nums := generateInt32s(100)
			SortInt32s(nums)
			if !IsSortedInt32s(nums) {
				t.Errorf("unsorted %b", nums)
			}
		})
	}
}

func BenchmarkSortUint32s(b *testing.B) {
	for i := 0; i < 1; i++ {
		b.Run(strconv.Itoa(i), func(b *testing.B) {
			nums := generateUint32s(b.N)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				SortUint32s(nums)
			}
		})
	}
}

func BenchmarkSortInt32s(b *testing.B) {
	for i := 0; i < 1; i++ {
		b.Run(strconv.Itoa(i), func(b *testing.B) {
			nums := generateInt32s(b.N)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				SortInt32s(nums)
			}
		})
	}
}

func generateUint32s(n int) []uint32 {
	nums := make([]uint32, n)
	for i := range nums {
		nums[i] = rand.Uint32()
	}
	return nums
}

func generateInt32s(n int) []int32 {
	nums := make([]int32, n)
	for i := range nums {
		nums[i] = int32(rand.Uint32())
	}
	return nums
}

func TestSortUint16s(t *testing.T) {
	for i := 0; i < 4; i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			nums := generateUint16s(100)
			SortUint16s(nums)
			if !IsSortedUint16s(nums) {
				t.Errorf("unsorted %b", nums)
			}
		})
	}
}

func TestSortInt16s(t *testing.T) {
	for i := 0; i < 4; i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			nums := generateInt16s(100)
			SortInt16s(nums)
			if !IsSortedInt16s(nums) {
				t.Errorf("unsorted %b", nums)
			}
		})
	}
}

func BenchmarkSortUint16s(b *testing.B) {
	for i := 0; i < 1; i++ {
		b.Run(strconv.Itoa(i), func(b *testing.B) {
			nums := generateUint16s(b.N)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				SortUint16s(nums)
			}
		})
	}
}

func BenchmarkSortInt16s(b *testing.B) {
	for i := 0; i < 1; i++ {
		b.Run(strconv.Itoa(i), func(b *testing.B) {
			nums := generateInt16s(b.N)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				SortInt16s(nums)
			}
		})
	}
}

func generateUint16s(n int) []uint16 {
	nums := make([]uint16, n)
	for i := range nums {
		nums[i] = uint16(rand.Uint32())
	}
	return nums
}

func generateInt16s(n int) []int16 {
	nums := make([]int16, n)
	for i := range nums {
		nums[i] = int16(rand.Uint32())
	}
	return nums
}

func TestSortUint8s(t *testing.T) {
	for i := 0; i < 4; i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			nums := generateUint8s(100)
			SortUint8s(nums)
			if !IsSortedUint8s(nums) {
				t.Errorf("unsorted %b", nums)
			}
		})
	}
}

func TestSortInt8s(t *testing.T) {
	for i := 0; i < 4; i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			nums := generateInt8s(100)
			SortInt8s(nums)
			if !IsSortedInt8s(nums) {
				t.Errorf("unsorted %b", nums)
			}
		})
	}
}

func BenchmarkSortUint8s(b *testing.B) {
	for i := 0; i < 1; i++ {
		b.Run(strconv.Itoa(i), func(b *testing.B) {
			nums := generateUint8s(b.N)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				SortUint8s(nums)
			}
		})
	}
}

func BenchmarkSortInt8s(b *testing.B) {
	for i := 0; i < 1; i++ {
		b.Run(strconv.Itoa(i), func(b *testing.B) {
			nums := generateInt8s(b.N)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				SortInt8s(nums)
			}
		})
	}
}

func generateUint8s(n int) []uint8 {
	nums := make([]uint8, n)
	for i := range nums {
		nums[i] = uint8(rand.Uint32())
	}
	return nums
}

func generateInt8s(n int) []int8 {
	nums := make([]int8, n)
	for i := range nums {
		nums[i] = int8(rand.Uint32())
	}
	return nums
}

func TestSortUints(t *testing.T) {
	for i := 0; i < 4; i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			nums := generateUints(100)
			SortUints(nums)
			if !IsSortedUints(nums) {
				t.Errorf("unsorted %b", nums)
			}
		})
	}
}

func TestSortInts(t *testing.T) {
	for i := 0; i < 4; i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			nums := generateInts(100)
			SortInts(nums)
			if !IsSortedInts(nums) {
				t.Errorf("unsorted %b", nums)
			}
		})
	}
}

func BenchmarkSortUints(b *testing.B) {
	for i := 0; i < 1; i++ {
		b.Run(strconv.Itoa(i), func(b *testing.B) {
			nums := generateUints(b.N)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				SortUints(nums)
			}
		})
	}
}

func BenchmarkSortInts(b *testing.B) {
	for i := 0; i < 1; i++ {
		b.Run(strconv.Itoa(i), func(b *testing.B) {
			nums := generateInts(b.N)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				SortInts(nums)
			}
		})
	}
}

func generateUints(n int) []uint {
	nums := make([]uint, n)
	for i := range nums {
		nums[i] = uint(rand.Uint64())
	}
	return nums
}

func generateInts(n int) []int {
	nums := make([]int, n)
	for i := range nums {
		nums[i] = int(rand.Uint64())
	}
	return nums
}
