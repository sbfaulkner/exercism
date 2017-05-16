package luhn

import (
	"strings"
	"unicode"
)

const testVersion = 2

// Valid determines whether or not a number is valid per the Luhn formula.
func Valid(number string) bool {
	number = reverse(strings.Replace(number, " ", "", -1))

	if len(number) <= 1 {
		return false
	}

	sum := 0

	for i, r := range number {
		if !unicode.IsDigit(r) {
			return false
		}

		digit := int(r - '0')
		if i%2 == 1 {
			digit = digit * 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
	}

	return sum%10 == 0
}

func reverse(input string) string {
	r := []rune(input)

	for i := 0; i < len(r)/2; i++ {
		r[i], r[len(r)-i-1] = r[len(r)-i-1], r[i]
	}

	return string(r)
}
