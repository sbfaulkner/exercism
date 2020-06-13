// Package grains provides methods to calculate the reward for the wise servant
package grains

import "errors"

// Square returns the number of grains on the specified square of the chessboard
func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, errors.New("invalid square - should be 1 <= n <= 64")
	}

	if n == 1 {
		return 1, nil
	}

	s, _ := Square(n - 1)

	return 2 * s, nil
}

// Total returns the total number of grains in the wise servant's reward
func Total() uint64 {
	t := uint64(0)

	for i := 1; i <= 64; i++ {
		s, _ := Square(i)
		t += s
	}

	return t
}
