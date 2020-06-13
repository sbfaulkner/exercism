package isbn

import (
	"unicode"
)

// IsValidISBN determines if a provided value is a valid ISBN-10
func IsValidISBN(input string) bool {
	var n int

	sum := 0
	digits := 0

	for _, x := range input {
		if x == '-' {
			continue
		}

		if unicode.IsDigit(x) {
			n = int(x - '0')
		} else if digits == 9 && x == 'X' {
			n = 10
		} else {
			return false
		}

		sum += n * (10 - digits)
		digits++
	}

	if digits != 10 {
		return false
	}

	return sum%11 == 0
}
