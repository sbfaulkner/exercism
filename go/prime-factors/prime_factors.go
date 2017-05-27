package prime

const testVersion = 2

// Factors returns the prime factors of a number.
func Factors(n int64) []int64 {
	factors := []int64{}

	f := int64(2)

	for n > 1 {
		if n%f == 0 {
			factors = append(factors, f)
			n /= f
		} else {
			f++
		}
	}

	return factors
}
