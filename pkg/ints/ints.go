package ints

import (
	"math"
	"math/bits"
)

const (
	// MaxUint defines the maximum value a uint can have
	MaxUint uint = (1 << bits.UintSize) - 1

	// MinUint defines the minimum value a uint can have
	MinUint uint = 0

	// MaxInt defines the maximum value an int can have
	MaxInt int = (1<<bits.UintSize)/2 - 1

	// MinInt defines the minimum value an int can have
	MinInt int = (1 << bits.UintSize) / -2

	// MaxUint8 defines the maximum value a uint8 can have
	MaxUint8 uint8 = math.MaxUint8

	// MinUint8 defines the minimum value a uint8 can have
	MinUint8 uint8 = 0

	// MaxInt8 defines the maximum value an int8 can have
	MaxInt8 int8 = math.MaxInt8

	// MinInt8 defines the minimum value an int8 can have
	MinInt8 int8 = math.MinInt8

	// MaxUint16 defines the maximum value a uint16 can have
	MaxUint16 uint16 = math.MaxUint16

	// MinUint16 defines the minimum value a uint16 can have
	MinUint16 uint16 = 0

	// MaxInt16 defines the maximum value an int16 can have
	MaxInt16 int16 = math.MaxInt16

	// MinInt16 defines the minimum value an int16 can have
	MinInt16 int16 = math.MinInt16

	// MaxUint32 defines the maximum value a uint32 can have
	MaxUint32 uint32 = math.MaxUint32

	// MinUint32 defines the minimum value a uint32 can have
	MinUint32 uint32 = 0

	// MaxInt32 defines the maximum value an int32 can have
	MaxInt32 int32 = math.MaxInt32

	// MinInt32 defines the minimum value an int32 can have
	MinInt32 int32 = math.MinInt32

	// MaxUint64 defines the maximum value a uint64 can have
	MaxUint64 uint64 = math.MaxUint64

	// MinUint64 defines the minimum value a uint64 can have
	MinUint64 uint64 = 0

	// MaxInt64 defines the maximum value an int64 can have
	MaxInt64 int64 = math.MaxInt64

	// MinInt64 defines the minimum value an int64 can have
	MinInt64 int64 = math.MinInt64
)

// UintMax returns the maximum among the given uint,
// or MinUint if no values are given.
func UintsMax(values ...uint) uint {
	value := MinUint
	for _, v := range values {
		if v > value {
			value = v
		}
	}
	return value
}

// UintMin returns the minimum among the given uint,
// or MaxUint if no values are given.
func UintsMin(values ...uint) uint {
	value := MaxUint
	for _, v := range values {
		if v < value {
			value = v
		}
	}
	return value
}

// IntMax returns the maximum among the given int,
// or MinInt if no values are given.
func IntsMax(values ...int) int {
	value := MinInt
	for _, v := range values {
		if v > value {
			value = v
		}
	}
	return value
}

// IntMin returns the minimum among the given int,
// or MaxInt if no values are given.
func IntsMin(values ...int) int {
	value := MaxInt
	for _, v := range values {
		if v < value {
			value = v
		}
	}
	return value
}

// Uint8Max returns the maximum among the given uint8,
// or MinUint8 if no values are given.
func Uint8sMax(values ...uint8) uint8 {
	value := MinUint8
	for _, v := range values {
		if v > value {
			value = v
		}
	}
	return value
}

// Uint8Min returns the minimum among the given uint8,
// or MaxUint8 if no values are given.
func Uint8sMin(values ...uint8) uint8 {
	value := MaxUint8
	for _, v := range values {
		if v < value {
			value = v
		}
	}
	return value
}

// Int8Max returns the maximum among the given int8,
// or MinInt8 if no values are given.
func Int8sMax(values ...int8) int8 {
	value := MinInt8
	for _, v := range values {
		if v > value {
			value = v
		}
	}
	return value
}

// Int8Min returns the minimum among the given int8,
// or MaxInt8 if no values are given.
func Int8sMin(values ...int8) int8 {
	value := MaxInt8
	for _, v := range values {
		if v < value {
			value = v
		}
	}
	return value
}

// Uint16Max returns the maximum among the given uint16,
// or MinUint16 if no values are given.
func Uint16sMax(values ...uint16) uint16 {
	value := MinUint16
	for _, v := range values {
		if v > value {
			value = v
		}
	}
	return value
}

// Uint16Min returns the minimum among the given uint16,
// or MaxUint16 if no values are given.
func Uint16sMin(values ...uint16) uint16 {
	value := MaxUint16
	for _, v := range values {
		if v < value {
			value = v
		}
	}
	return value
}

// Int16Max returns the maximum among the given int16,
// or MinInt16 if no values are given.
func Int16sMax(values ...int16) int16 {
	value := MinInt16
	for _, v := range values {
		if v > value {
			value = v
		}
	}
	return value
}

// Int16Min returns the minimum among the given int16,
// or MaxInt16 if no values are given.
func Int16sMin(values ...int16) int16 {
	value := MaxInt16
	for _, v := range values {
		if v < value {
			value = v
		}
	}
	return value
}

// Uint32Max returns the maximum among the given uint32,
// or MinUint32 if no values are given.
func Uint32sMax(values ...uint32) uint32 {
	value := MinUint32
	for _, v := range values {
		if v > value {
			value = v
		}
	}
	return value
}

// Uint32Min returns the minimum among the given uint32,
// or MaxUint32 if no values are given.
func Uint32sMin(values ...uint32) uint32 {
	value := MaxUint32
	for _, v := range values {
		if v < value {
			value = v
		}
	}
	return value
}

// Int32Max returns the maximum among the given int32,
// or MinInt32 if no values are given.
func Int32sMax(values ...int32) int32 {
	value := MinInt32
	for _, v := range values {
		if v > value {
			value = v
		}
	}
	return value
}

// Int32Min returns the minimum among the given int32,
// or MaxInt32 if no values are given.
func Int32sMin(values ...int32) int32 {
	value := MaxInt32
	for _, v := range values {
		if v < value {
			value = v
		}
	}
	return value
}

// Uint64Max returns the maximum among the given uint64,
// or MinUint64 if no values are given.
func Uint64sMax(values ...uint64) uint64 {
	value := MinUint64
	for _, v := range values {
		if v > value {
			value = v
		}
	}
	return value
}

// Uint64Min returns the minimum among the given uint64,
// or MaxUint64 if no values are given.
func Uint64sMin(values ...uint64) uint64 {
	value := MaxUint64
	for _, v := range values {
		if v < value {
			value = v
		}
	}
	return value
}

// Int64Max returns the maximum among the given int64,
// or MinInt64 if no values are given.
func Int64sMax(values ...int64) int64 {
	value := MinInt64
	for _, v := range values {
		if v > value {
			value = v
		}
	}
	return value
}

// Int64Min returns the minimum among the given int64,
// or MaxInt64 if no values are given.
func Int64sMin(values ...int64) int64 {
	value := MaxInt64
	for _, v := range values {
		if v < value {
			value = v
		}
	}
	return value
}
