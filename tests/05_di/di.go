package main

import (
	"bytes"
	"fmt"
)

// Greet returns hello plus argument
func Greet(writer *bytes.Buffer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func main() {
	//Greet("James")
}
