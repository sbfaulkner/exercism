package isogram

import (
	"strings"
	"unicode"
)

const testVersion = 1

// IsIsogram returns a boolean indicating whether or not a word or phrase is an isogram.
func IsIsogram(text string) bool {
	letters := map[rune]bool{}

	for _, letter := range strings.ToLower(text) {
		if !unicode.IsLetter(letter) {
			continue
		}
		if letters[letter] {
			return false
		}
		letters[letter] = true
	}

	return true
}
