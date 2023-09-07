package nlp_test

import (
	"fmt"
	"samples-go/practical-go/nlp"
)

func ExampleTokenize() {
	text := "Hello there!"
	tokens := nlp.Tokenize(text)
	fmt.Println(tokens)

	//Output:
	// [hello there]
}

/*
Test discovery:
For every file ending with _test.go, run every function that matches either:

- Example[A-Z_].*
- Test[A-Z_].*
*/
