package allyourbase

import (
	"errors"
)

const testVersion = 1

// common errors
var (
	ErrInvalidDigit = errors.New("allyourbase: invalid digit")
	ErrInvalidBase  = errors.New("allyourbase: invalid base")
)

func ConvertToBase(inputBase uint64, inputDigits []uint64, outputBase uint64) ([]uint64, error) {
	if inputBase < 2 || outputBase < 2 {
		return []uint64{}, ErrInvalidBase
	}

	n, err := fromBase(inputBase, inputDigits)

	if err != nil {
		return []uint64{}, err
	}

	if n == 0 {
		return []uint64{0}, nil
	}

	return toBase(outputBase, n), nil
}

func fromBase(base uint64, digits []uint64) (uint64, error) {
	if len(digits) == 0 {
		return 0, nil
	}

	if len(digits) == 1 {
		return digits[0], nil
	}

	d := digits[len(digits)-1]

	if d >= base {
		return 0, ErrInvalidDigit
	}

	n, err := fromBase(base, digits[0:len(digits)-1])

	if err != nil {
		return 0, err
	}

	return n*base + d, nil
}

func toBase(base uint64, n uint64) []uint64 {
	if n == 0 {
		return []uint64{}
	}

	return append(toBase(base, n/base), n%base)
}
