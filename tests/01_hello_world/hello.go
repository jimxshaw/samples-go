package main

import (
	"fmt"
)

const helloPrefix = "Hello, "

// Hello prints out a string.
func Hello(name string) string {

	if len(name) > 0 {
		return helloPrefix + name
	}

	return helloPrefix + "World"
}

func main() {
	fmt.Println(Hello("James"))
}
