package nlp

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTokenize(t *testing.T) {
	text := "Hello there!"
	expected := []string{"hello", "there"}
	tokens := Tokenize(text)

	require.Equal(t, expected, tokens)

	// Before using testify:
	// if tokens != expected { // Cannot compare slices with == in Go (only to nil)
	// if !reflect.DeepEqual(expected, tokens) {
	// 	t.Fatalf("expected %#v, got %#v", expected, tokens)
	// }
}
