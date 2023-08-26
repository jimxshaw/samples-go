package nlp

import (
	"regexp"
	"strings"
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
		tokens = append(tokens, token)
	}
	return tokens
}
