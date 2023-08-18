package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("goroutine")
	fmt.Println("main")

	for i := 0; i < 3; i++ {
		// BUG: all goroutines use the same "i" for the for loop
		// A closure bug, not a goroutine bug.
		// go func() {
		// 	fmt.Println(i) // for loop i
		// }()

		// Fix 1: use a parameter
		// go func(n int) {
		// 	fmt.Println(n)
		// }(i)

		// Fix 2: use a loop body variable
		i := i // this i "shadows" the i for the for loop
		go func() {
			fmt.Println(i) // body variable i
		}()
	}

	time.Sleep(10 * time.Millisecond)

	shadowExample()
}

func shadowExample() {
	n := 4
	{
		n := 2 // from here n is the inner n, not the outer
		fmt.Println("inner: ", n)
	}
	fmt.Println("outer: ", n)
}
