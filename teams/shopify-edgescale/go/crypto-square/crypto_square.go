package cryptosquare

import (
	"strings"
	"unicode"
)

func normalize(text string) []rune {
	n := make([]rune, 0, len(text))

	for _, r := range text {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			continue
		}

		n = append(n, unicode.ToLower(r))
	}

	return n
}

func dimensions(text []rune) (int, int) {
	r := 0
	c := 0

	for {
		if c*r >= len(text) {
			break
		}

		c++

		if c*r >= len(text) {
			break
		}

		r++
	}

	return r, c
}

// Encode encrypts a secret message using square code
func Encode(input string) string {
	normalized := normalize(input)

	rows, columns := dimensions(normalized)

	square := make([][]rune, columns)

	for c := 0; c < columns; c++ {
		square[c] = make([]rune, rows)

		for r := 0; r < rows; r++ {
			i := r*columns + c
			if i >= len(normalized) {
				square[c][r] = ' '
			} else {
				square[c][r] = normalized[i]
			}
		}
	}

	s := make([]string, 0, columns)
	for _, c := range square {
		s = append(s, string(c))
	}

	return strings.Join(s, " ")
}
