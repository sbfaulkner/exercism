package allyourbase

import (
	"errors"
)

func fromBase(base int, in []int) (int, error) {
	if base < 2 {
		return 0, errors.New("input base must be >= 2")
	}

	n := 0

	for _, i := range in {
		if i < 0 || i >= base {
			return 0, errors.New("all digits must satisfy 0 <= d < input base")
		}

		n *= base
		n += i
	}

	return n, nil
}

func toBase(base int, n int) ([]int, error) {
	if base < 2 {
		return nil, errors.New("output base must be >= 2")
	}

	if n == 0 {
		return []int{0}, nil
	}

	o := make([]int, 0, 1)

	for n > 0 {
		o = append([]int{n % base}, o...)
		n /= base
	}

	return o, nil
}

// ConvertToBase converts a number, represented as a sequence of digits in one base, to any other base
func ConvertToBase(ibase int, in []int, obase int) ([]int, error) {
	n, err := fromBase(ibase, in)
	if err != nil {
		return nil, err
	}

	return toBase(obase, n)
}
