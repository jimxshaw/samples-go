package main

// Build a container from scratch.
// docker					run image <cmd> <params>
// go run main.go run				<cmd> <params>

func must(err error) {
	if err != nil {
		panic(err)
	}
}
