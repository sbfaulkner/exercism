package lsproduct

import "errors"
import "unicode"
import "fmt"

const testVersion = 5

// LargestSeriesProduct calculates the largest product for a contiguous substring of digits of length n
// given a string of digits.
func LargestSeriesProduct(stringOfDigits string, n int) (int, error) {
	if n > len(stringOfDigits) {
		return 0, errors.New("insufficient digits")
	}

	if n < 0 {
		return 0, errors.New("invalid number of digits specified")
	}

	count := len(stringOfDigits) - n + 1
	largest := 0

	for i := 0; i < count; i++ {
		product, err := seriesProduct(stringOfDigits[i : i+n])
		if err != nil {
			return 0, err
		}

		if product > largest {
			largest = product
		}
	}

	return largest, nil
}

func seriesProduct(stringOfDigits string) (int, error) {
	product := 1

	for _, r := range stringOfDigits {
		if !unicode.IsDigit(r) {
			return 0, fmt.Errorf("not a digit - %q", r)
		}
		product *= int(r - '0')
	}

	return product, nil
}
