package pangram

import (
	"strings"
)

const testVersion = 1

// IsPangram returns true if the sentence is a pangram
func IsPangram(sentence string) bool {
	runes := make(map[rune]bool)

	for _, ch := range strings.ToLower(sentence) {
		if 'a' <= ch && ch <= 'z' {
			runes[ch] = true
		}
	}

	return len(runes) == 26
}
