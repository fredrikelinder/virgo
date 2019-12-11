package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	n, _ := strconv.ParseUint(os.Args[1], 10, 64)
	before := time.Now()
	ints.Prime(n)
	duration := time.Since(before)
	fmt.Println(duration)
}
