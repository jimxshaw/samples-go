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

	prefix := helloPrefix

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case chinese:
		prefix = chineseHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("James", chinese))
}
