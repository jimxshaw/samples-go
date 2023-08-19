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
		go func(n int) {
			fmt.Println(n)
		}(i)

		// Fix 2: use a loop body variable
		// i := i // this i "shadows" the i for the for loop
		// go func() {
		// 	fmt.Println(i) // body variable i
		// }()

		// Fake Fix 3: put the goroutine somewhere else, do not do this
		// func() {
		// 	go fmt.Println(i)
		// }()

		// Fake Fix 4: remove the anonymous function, do not do this
		// go fmt.Println(i)
	}

	time.Sleep(10 * time.Millisecond)

	shadowExample()

	// Channel only has 2 operations
	ch := make(chan string)
	go func() {
		// Think of a channel as a conveyor belt where a worker, goroutine,
		// puts an item on the conveyor, in this case a string.
		ch <- "hello!" // Send operation
	}()
	// The main goroutine is there to receive the item from the conveyor.
	msg := <-ch // Receive operation

	fmt.Println(msg)

	go func() {
		for i := 0; i < 3; i++ {
			msg := fmt.Sprintf("message #%d", i+1)
			ch <- msg
		}
		close(ch)
	}()

	for msg := range ch {
		fmt.Println("got: ", msg)
	}

	// The for range loop over a channel essentially does this.
	// for {
	// 	msg, ok := <-ch
	// 	if !ok {
	// 		break
	// 	}
	// 	fmt.Println("got: ", msg)
	// }

	result := <-ch // closed channel
	fmt.Printf("closed: %#v\n", result)

	result, ok := <-ch // closed channel
	fmt.Printf("closed: %#v (ok = %v)\n", result, ok)

	// ch <- "another message" // If ch is closed then it will panic.
}

// Channel Semantics
// Send & Receive will block until opposite operation (*).
// Receive from a closed channel will return the zero value without blocking.
// Send to a closed channel will panic.
// Closing a closed channel will also panic.

func shadowExample() {
	n := 4
	{ // Naked { } Block that only introduces a new lexical scope, rarely used syntax.
		n := 2 // from here n is the inner n, not the outer
		fmt.Println("inner: ", n)
	}
	fmt.Println("outer: ", n)
}
