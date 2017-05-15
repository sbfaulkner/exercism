package grains

import "fmt"

const testVersion = 1

// Square returns the number of grains of wheat on a given square of the chessboard
func Square(n int) (count uint64, e error) {
	if n < 1 || n > 64 {
		e = fmt.Errorf("invalid square - %d", n)
	} else {
		count = 1 << uint(n-1)
	}

	return
}

// Total returns the total number of grains of wheat on the chessboard
func Total() (total uint64) {
	for i := 1; i <= 64; i++ {
		count, _ := Square(i)
		total += count
	}

	return
}
