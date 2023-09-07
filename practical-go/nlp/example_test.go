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
