package main

import (
	"fmt"
	"os"
	"os/exec"
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

	// Run a command and its arguments.
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Run()
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
