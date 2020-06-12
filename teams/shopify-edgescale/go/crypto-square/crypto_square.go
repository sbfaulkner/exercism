package cryptosquare

import (
	"fmt"
	"strings"
	"unicode"
)

func normalize(r rune) rune {
	if unicode.IsUpper(r) {
		return unicode.ToLower(r)
	}

	if unicode.IsLower(r) || unicode.IsDigit(r) {
		return r
	}

	return -1
}

// Encode encrypts a secret message using square code
func Encode(input string) string {
	normalized := strings.Map(normalize, input)

	length := len(normalized)

	rows := 0
	columns := 0

	for {
		if columns*rows >= length {
			break
		}

		columns++

		if columns*rows >= length {
			break
		}

		rows++
	}

	square := make([]string, columns)

	for i, r := range []rune(fmt.Sprintf("%-*s", rows*columns, normalized)) {
		cc := i % columns

		square[cc] += string(r)
	}

	return strings.Join(square, " ")
}
