package main

import (
	"fmt"
)

const helloPrefix = "Hello, "

// Hello prints out a string.
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	if language == "Spanish" {
		return "Hola, " + name
	}

	return helloPrefix + name
}

func main() {
	fmt.Println(Hello("James", "Spanish"))
}
