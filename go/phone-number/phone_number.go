package phonenumber

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

const testVersion = 2

// ErrPhoneNumberInvalid is returned when a NANP phone number is not able to be sanitized.
var ErrPhoneNumberInvalid = errors.New("phonenumber: invalid phone number")

// regular expression for digits
var phoneNumberFormat = regexp.MustCompile("\\A1?(?:[2-9][0-9]{2}){2}[0-9]{4}\\z")

// AreaCode extracts the area code from a NANP phone number after sanitizing.
func AreaCode(input string) (string, error) {
	number, err := Number(input)

	if err != nil {
		return "", err
	}

	return number[0:3], nil
}

// Format sanitizes and formats a NANP phone number.
func Format(input string) (string, error) {
	number, err := Number(input)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("(%s) %s-%s", number[0:3], number[3:6], number[6:10]), nil
}

// Number sanitizes a NANP phone number.
func Number(input string) (string, error) {
	d := digits(input)

	if !phoneNumberFormat.MatchString(d) {
		return "", ErrPhoneNumberInvalid
	}

	return d[len(d)-10 : len(d)], nil
}

func digits(input string) string {
	return strings.Map(mapDigit, input)
}

func mapDigit(r rune) rune {
	if !unicode.IsOneOf([]*unicode.RangeTable{unicode.Digit, unicode.Letter}, r) {
		return -1
	}
	return r
}
