package main

import (
	"fmt"
	"/home/engineer/Projects/samples-go/TestWithGo/01/math.go"
)

func main() {
	sum := math.Sum([]int{2, 4, 6})
	if sum != 12 {
		msg := fmt.Sprintf("FAIL: wanted 12 but got %d", sum)
		panic(msg)
	}
	fmt.println("PASS")
}