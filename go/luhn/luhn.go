package luhn

import (
	"fmt"
	"strings"
)

const testVersion = 2

// Valid determines whether or not a number is valid per the Luhn formula.
func Valid(number string) bool {
	number = strings.Replace(number, " ", "", -1)

	sum := 0
	index := len(number) - 1

	if index < 1 {
		return false
	}

	for i := range number {
		digit := int(number[index] - '0')

		fmt.Println(digit)
		if i%2 == 1 {
			digit = digit * 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
		index--
	}

	return sum%10 == 0
}
