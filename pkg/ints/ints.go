package ints

const (
	// MaxUint defines the maximum value a uint can have
	MaxUint = ^uint(0)

	// MinUint defines the minimum value a uint can have
	MinUint = 0

	// MaxInt defines the maximum value an int can have
	MaxInt = int(^uint(0) >> 1)

	// MinUint defines the minimum value an int can have
	MinInt = -MaxInt - 1
)

// Max returns the maximum value among the given ints.
func Max(value int, values ...int) int {
	for _, v := range values {
		if v > value {
			value = v
		}
	}

	return value
}

// Max returns the minimum value among the given ints.
func Min(value int, values ...int) int {
	for _, v := range values {
		if v < value {
			value = v
		}
	}

	return value
}
