package main

import (
	"fmt"
	"io"
	"os"
)

// Greet returns hello plus argument
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func main() {
	Greet(os.Stdout, "James")
}
