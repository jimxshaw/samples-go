package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
}

func main() {
	luke := person{
		firstName: "Luke",
		lastName:  "Skywalker"}

	fmt.Println(luke)
}
