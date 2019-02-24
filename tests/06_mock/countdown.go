package main

import (
	"fmt"
	"io"
	"os"
)

// Countdown counts down
func Countdown(out io.Writer) {
	fmt.Fprint(out, "3")
}

func main() {
	Countdown(os.Stdout)
}
