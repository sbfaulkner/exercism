package atbash

import (
	"strings"
	"unicode"
)

const testVersion = 2

// Atbash uses an atbash cipher to encrypt a string.
func Atbash(input string) string {
	return strings.Join(slice(strings.Map(replacement, strings.ToLower(input)), 5), " ")
}

func replacement(r rune) rune {
	if unicode.IsDigit(r) {
		return r
	}

	if !unicode.IsLetter(r) {
		return -1
	}

	return 'a' + 'z' - r
}

func slice(input string, l int) []string {
	slices := make([]string, (len(input)+l-1)/l)

	for i := range slices {
		j := i*l + l

		if j > len(input) {
			j = len(input)
		}

		slices[i] = input[i*l : j]
	}

	return slices
}
