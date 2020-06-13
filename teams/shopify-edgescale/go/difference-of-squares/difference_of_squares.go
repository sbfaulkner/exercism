package diffsquares

// SquareOfSum returns the square of the sum of the first N natural numbers
func SquareOfSum(n int) int {
	s := 0

	for n > 0 {
		s += n
		n--
	}

	return s * s
}

// SumOfSquares returns the sum of the squares of the first N natural numbers
func SumOfSquares(n int) int {
	s := 0

	for n > 0 {
		s += (n * n)
		n--
	}

	return s
}

// Difference returns the difference between the square of the sum and the sum of the squares of the first N natural numbers
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
