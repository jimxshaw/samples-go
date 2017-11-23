package main

import "fmt"

func main() {
	mySlice := []string{"Hello", "How", "Are", "You"}

	updateSlice(mySlice)

	fmt.Println(mySlice)

	// In Go, whenever a slice data structure is created,
	// an array of which that slice points to is also created.
	// A slice is still pass by value, meaning a copy is created,
	// but the underlying array that connects with the slice is NOT
	// copied and the array is the structure that holds values.
	// Hence why the "slice" is modifed upon the updateSlice
	// function call. Because arrays are rarely ever used and
	// slices are very frequently used, a slice structure is
	// treated as a reference type.

	// Value Types: int, float, string, bool, structs
	// Reference Types: slices, maps, channels, pointers, functions
}

func updateSlice(s []string) {
	s[0] = "Bye"
}
