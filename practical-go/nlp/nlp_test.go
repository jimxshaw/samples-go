package nlp

import (
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/stretchr/testify/require"
)

// var tokenizeCases = []struct { // anonymous struct
// 	text   string
// 	tokens []string
// }{
// 	{"Hello there!", []string{"hello", "there"}},
// 	{"", nil},
// }

// Read test cases from tokenize_cases.toml.
type tokenizeCase struct {
	Text   string
	Tokens []string
}

func loadTokenizeCases(t *testing.T) []tokenizeCase {
	var testCases struct {
		Cases []tokenizeCase
	}

	// Decode by reading the file then use Unmarshal.
	// data, err := os.ReadFile("tokenize_cases.toml")
	// require.NoError(t, err, "Read file")
	// toml.Unmarshal(data, &testCases)

	// Decode the file directly with DecodeFile.
	_, err := toml.DecodeFile("tokenize_cases.toml", &testCases)
	require.NoError(t, err, "Unmarshal TOML")

	return testCases.Cases
}

// This is called regression testing or table testing.
func TestTokenizeTable(t *testing.T) {
	// for _, tc := range tokenizeCases {
	for _, tc := range loadTokenizeCases(t) {
		t.Run(tc.Text, func(t *testing.T) {
			tokens := Tokenize(tc.Text)
			require.Equal(t, tc.Tokens, tokens)
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
