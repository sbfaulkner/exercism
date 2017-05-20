package romannumerals

import (
	"fmt"
	"strings"
)

const testVersion = 3

var tens = "IXCM"
var fives = "VLD"

// ToRomanNumeral converts a number to roman numerals.
func ToRomanNumeral(number int) (string, error) {
	var roman string

	fmt.Println(number)
	if number <= 0 {
		return "", fmt.Errorf("number must be positive")
	}

	if number >= 4000 {
		return "", fmt.Errorf("number must be less than 4000")
	}

	for i := 0; i < len(tens); i++ {
		switch digit := number % 10; digit {
		case 1, 2, 3:
			roman = strings.Repeat(tens[i:i+1], digit) + roman
		case 4:
			roman = tens[i:i+1] + fives[i:i+1] + roman
		case 5:
			roman = fives[i:i+1] + roman
		case 6, 7, 8:
			roman = fives[i:i+1] + strings.Repeat(tens[i:i+1], digit-5) + roman
		case 9:
			roman = tens[i:i+2] + roman
		}

		number /= 10
	}

	return roman, nil
}
