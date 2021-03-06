package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Shaw",
		contact: contactInfo{
			email:   "jimxshaw@gmail.com",
			zipCode: 90210,
		},
	}

	// Turn memory address into value with *address.
	// Turn value into memory address with &value.

	// The ampersand & is an operator that says give
	// the memory address of the value to which this
	// variable is pointing.
	//jimPointer := &jim
	//jimPointer.updateName("James")

	// This is a shortcut that does the exact same
	// thing as the above two lines.
	// In Go, if we define a receiver function with a pointer
	// to a type, Go will allow either that pointer
	// type or that type itself to call the function.
	jim.updateName("Jimmy")
	jim.print()

}

// An asterisk * in front of a type means a type description.
// It says we're working with a pointer to a person type.
func (pointerToPerson *person) updateName(newFirstName string) {
	// An asterisk * in front of an actual pointer means we
	// want to manipulate the value that pointer is referencing.
	(*pointerToPerson).firstName = newFirstName
}

func (pointerToPerson *person) print() {
	fmt.Printf("%+v", (*pointerToPerson))
}
