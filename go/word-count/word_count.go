package wordcount

import (
	"strings"
	"unicode"
)

const testVersion = 3

// Frequency is used to represent the frequency of words in a phrase.
type Frequency map[string]int

// WordCount counts the frequency of each word in the phrase.
func WordCount(phrase string) Frequency {
	counts := make(Frequency)

	for _, w := range strings.FieldsFunc(normalize(phrase), wordBoundary) {
		counts[w]++
	}

	return counts
}

func normalize(s string) string {
	s = strings.Replace(s, " '", " ", -1)
	s = strings.Replace(s, "' ", " ", -1)
	return strings.ToLower(s)
}

func wordBoundary(r rune) bool {
	return unicode.In(r, unicode.Punct, unicode.White_Space, unicode.Symbol) && r != '\''
}
