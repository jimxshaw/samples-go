package nlp

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var tokenizeCases = []struct { // anonymous struct
	text   string
	tokens []string
}{
	{"Hello there!", []string{"hello", "there"}},
	{"", nil},
}

// This is called regression testing or table testing.
func TestTokenizeTable(t *testing.T) {
	for _, tc := range tokenizeCases {
		t.Run(tc.text, func(t *testing.T) {
			tokens := Tokenize(tc.text)
			require.Equal(t, tc.tokens, tokens)
		})
	}
}

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
