package isogram

import (
	"unicode"
)

func ignoredRune(r rune) bool {
	return !unicode.IsLetter(r)
}

// IsIsogram determines if a word or phrase is an isogram
func IsIsogram(text string) bool {
	letters := make(map[rune]bool, len(text))

	for _, r := range text {
		if ignoredRune(r) {
			continue
		}

		r = unicode.ToLower(r)

		if letters[r] {
			return false
		}

		letters[r] = true
	}

	return true
}
