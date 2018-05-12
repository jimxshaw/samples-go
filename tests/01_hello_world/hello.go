package main

import (
	"fmt"
)

const helloPrefix = "Hello, "

// Hello prints out a string.
func Hello(name string) string {
	if name == "" {
		name = "James"
	}

	return helloPrefix + name
}

func main() {
	fmt.Println(Hello("James"))
}
