package acronym

import (
	"strings"
	"unicode"
)

const testVersion = 2

// Abbreviate converts an input string into the representative acronym
func Abbreviate(value string) (acronym string) {
	isDelimiter := func(c rune) bool {
		return !unicode.IsLetter(c)
	}

	words := strings.FieldsFunc(value, isDelimiter)

	for _, word := range words {
		acronym += strings.ToUpper(word[0:1])
	}

	return
}
