package main

import (
	"fmt"
	"log"
)

func main() {
	// fmt.Println(divide(1, 0))
	fmt.Println(safeDivide(1, 0))
}

// named return values
func safeDivide(a, b int) (q int, err error) {
	// q and err are local variables inside safeDivide
	defer func() {
		if e := recover(); e != nil {
			log.Println("ERROR: ", e)
			err = fmt.Errorf("%v", e)
		}
	}()

	return a / b, nil
	/* It's possible to do this in Go but frowned upon
	q = a / b
	return
	*/
}

func divide(a, b int) int {
	return a / b
}
