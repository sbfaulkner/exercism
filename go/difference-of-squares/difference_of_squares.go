package diffsquares

const testVersion = 1

// SquareOfSums calculates the square of the sum of the first `count` terms
func SquareOfSums(count int) (value int) {
	for term := 1; term <= count; term++ {
		value += term
	}

	value *= value

	return
}

// SumOfSquares calculates the sum of the squares of the first `count` terms
func SumOfSquares(count int) (value int) {
	for term := 1; term <= count; term++ {
		value += term * term
	}

	return
}

// Difference calculates the difference between the square of the sum and the sum of the squares
func Difference(count int) (difference int) {
	return SquareOfSums(count) - SumOfSquares(count)
}
