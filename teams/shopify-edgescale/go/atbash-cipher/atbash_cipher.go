package atbash

import (
	"unicode"
)

const groupLength = 5

func atbash(r rune) rune {
	switch {
	case unicode.IsLower(r):
		return 25 - (r - 'a') + 'a'

	case unicode.IsUpper(r):
		return 25 - (r - 'A') + 'a'

	case unicode.IsDigit(r):
		return r
	}

	return -1
}

// Atbash encrypts a string using the Atbash cipher
func Atbash(s string) string {
	o := make([]rune, 0, len(s)+len(s)/5)
	l := 0

	for _, r := range s {
		a := atbash(r)

		if a > 0 {
			if l%groupLength == 0 && l > 0 {
				o = append(o, ' ')
			}

			o = append(o, a)
			l++
		}
	}

	return string(o)
}
