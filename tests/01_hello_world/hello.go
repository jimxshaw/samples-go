package main

import (
	"fmt"
)

const spanish = "Spanish"
const french = "French"
const chinese = "Chinese"
const helloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const chineseHelloPrefix = "Ni hao, "

// Hello prints out a string.
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	// Having a named return value means a variable
	// is declared by that name and assigned it's
	// default value. Calling return is enough when
	// using a named return value in the method signature.
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case chinese:
		prefix = chineseHelloPrefix
	default:
		prefix = helloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("James", chinese))
}
