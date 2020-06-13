package isogram

import "strings"

const ignored = "- "

func ignoredRune(r rune) bool {
	return strings.ContainsRune(ignored, r)
}

// IsIsogram determines if a word or phrase is an isogram
func IsIsogram(text string) bool {
	letters := map[rune]bool{}

	for _, r := range strings.ToLower(text) {
		if ignoredRune(r) {
			continue
		}

		if letters[r] {
			return false
		}

		letters[r] = true
	}

	return true
}
