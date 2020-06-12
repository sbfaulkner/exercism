package cryptosquare

import (
	"fmt"
	"strings"
	"unicode"
)

func normalize(r rune) rune {
	if unicode.IsLower(r) || unicode.IsDigit(r) {
		return r
	}

	if unicode.IsUpper(r) {
		return unicode.ToLower(r)
	}

	return -1
}

func dimensions(s string) (int, int) {
	r := 0
	c := 0

	for {
		if c*r >= len(s) {
			break
		}

		c++

		if c*r >= len(s) {
			break
		}

		r++
	}

	return r, c
}

// Encode encrypts a secret message using square code
func Encode(input string) string {
	normalized := strings.Map(normalize, input)

	rows, columns := dimensions(normalized)

	square := make([][]rune, columns)

	for c := range square {
		square[c] = make([]rune, rows)
	}

	for i, r := range []rune(fmt.Sprintf("%-*s", rows*columns, normalized)) {
		square[i%columns][i/columns] = r
	}

	s := make([]string, 0, columns)
	for _, c := range square {
		s = append(s, string(c))
	}

	return strings.Join(s, " ")
}
