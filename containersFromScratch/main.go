package main

import (
	"fmt"
	"os"
)

// Build a container from scratch.
// docker					run image <cmd> <params>
// go run main.go run				<cmd> <params>
func main() {
	switch os.Args[1] {
	case "run":
		run()
	default:
		panic("Bad Command!")
	}
}

func run() {
	fmt.Printf("Running %v\n", os.Args[2:])
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
