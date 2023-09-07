package nlp

import (
	"reflect"
	"testing"
)

func TestTokenize(t *testing.T) {
	text := "Hello there!"
	expected := []string{"hello", "there"}
	tokens := Tokenize(text)

	// if tokens != expected { // Cannot compare slices with == in Go (only to nil)
	if !reflect.DeepEqual(expected, tokens) {
		t.Fatalf("expected %#v, got %#v", expected, tokens)
	}
}
