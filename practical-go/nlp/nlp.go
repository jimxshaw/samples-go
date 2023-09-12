package nlp

import (
	"regexp"
	"strings"

	"samples-go/practical-go/nlp/stemmer"
)

var (
	wordRe = regexp.MustCompile(`[a-zA-Z]+`)
)

// Tokenize returns a list of (lower case) tokens in the text.
func Tokenize(text string) []string {
	words := wordRe.FindAllString(text, -1)
	var tokens []string
	for _, w := range words {
		token := strings.ToLower(w)
		token = stemmer.Stem(token)

		if len(token) != 0 {
			tokens = append(tokens, token)
		}
	}
	return tokens
}
