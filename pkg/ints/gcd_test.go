package ints

import "fmt"

func ExampleUint64GCD() {
	fmt.Println("gcd:", Uint64GCD(12, 15))
	// Output:
	// gcd: 3
}

func ExampleInt64GCD() {
	fmt.Println("gcd:", Int64GCD(12, 15))
	fmt.Println("gcd:", Int64GCD(12, -15))
	fmt.Println("gcd:", Int64GCD(-12, 15))
	fmt.Println("gcd:", Int64GCD(-12, -15))
	// Output:
	// gcd: 3
	// gcd: 3
	// gcd: 3
	// gcd: 3
}

func ExampleUint32GCD() {
	fmt.Println(Uint32GCD(12, 15))
	// Output: 3
}

func ExampleInt32GCD() {
	fmt.Println(Int32GCD(12, 15))
	// Output: 3
}

func ExampleUint16GCD() {
	fmt.Println(Uint16GCD(12, 15))
	// Output: 3
}

func ExampleInt16GCD() {
	fmt.Println(Int16GCD(12, 15))
	// Output: 3
}

func ExampleUint8GCD() {
	fmt.Println(Uint8GCD(12, 15))
	// Output: 3
}

func ExampleInt8GCD() {
	fmt.Println(Int8GCD(12, 15))
	// Output: 3
}

func ExampleUintGCD() {
	fmt.Println(UintGCD(12, 15))
	// Output: 3
}

func ExampleIntGCD() {
	fmt.Println(IntGCD(12, 15))
	// Output: 3
}
