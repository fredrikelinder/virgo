package ints

// Uint64GCD returns the greatest common divisor
// between the two given numbers.
func Uint64GCD(a, b uint64) uint64 {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

// Int64GCD returns the greatest common divisor
// between the two given numbers.
func Int64GCD(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}

	if a < 0 {
		return -a
	}

	return a
}

// Uint32GCD returns the greatest common divisor
// between the two given numbers.
func Uint32GCD(a, b uint32) uint32 {
	return uint32(Uint64GCD(uint64(a), uint64(b)))
}

// Int32GCD returns the greatest common divisor
// between the two given numbers.
func Int32GCD(a, b int32) int32 {
	return int32(Int64GCD(int64(a), int64(b)))
}

// Uint16GCD returns the greatest common divisor
// between the two given numbers.
func Uint16GCD(a, b uint16) uint16 {
	return uint16(Uint64GCD(uint64(a), uint64(b)))
}

// Int16GCD returns the greatest common divisor
// between the two given numbers.
func Int16GCD(a, b int16) int16 {
	return int16(Int64GCD(int64(a), int64(b)))
}

// Uint8GCD returns the greatest common divisor
// between the two given numbers.
func Uint8GCD(a, b uint8) uint8 {
	return uint8(Uint64GCD(uint64(a), uint64(b)))
}

// Int8GCD returns the greatest common divisor
// between the two given numbers.
func Int8GCD(a, b int8) int8 {
	return int8(Int64GCD(int64(a), int64(b)))
}

// UintGCD returns the greatest common divisor
// between the two given numbers.
func UintGCD(a, b uint) uint {
	return uint(Uint64GCD(uint64(a), uint64(b)))
}

// IntGCD returns the greatest common divisor
// between the two given numbers.
func IntGCD(a, b int) int {
	return int(Int64GCD(int64(a), int64(b)))
}
