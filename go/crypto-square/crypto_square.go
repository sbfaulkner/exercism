package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

const testVersion = 2

// Encode uses the classic square code method to encrypt a message.
func Encode(text string) string {
	normalize := func(r rune) rune {
		if !unicode.In(r, unicode.Letter, unicode.Number) {
			return -1
		}
		return unicode.ToLower(r)
	}

	text = strings.Map(normalize, text)
	columns := int(math.Ceil(math.Sqrt(float64(len(text)))))
	square := make([]string, columns)

	for c := 0; c < columns; c++ {
		for r := c; r < len(text); r += columns {
			square[c] += text[r : r+1]
		}
	}

	return strings.Join(square, " ")
}
