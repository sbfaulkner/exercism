// Package grains provides methods to calculate the reward for the wise servant
package grains

import "errors"

// Square returns the number of grains on the specified square of the chessboard
func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, errors.New("invalid square - should be 1 <= n <= 64")
	}

	return 1 << (n - 1), nil
}

// Total returns the total number of grains in the wise servant's reward
func Total() uint64 {
	total := uint64(0)

	for i := 1; i <= 64; i++ {
		grains, _ := Square(i)
		total += grains
	}

	return total
}
