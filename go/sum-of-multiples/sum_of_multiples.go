package summultiples

const testVersion = 1

// SumMultiples calculates the sum of all multiples of the provided divisors less than the specified limit
func SumMultiples(limit int, divisors ...int) (sum int) {
	multiples := map[int][]int{}

	for _, divisor := range divisors {
		for n := divisor; n < limit; n += divisor {
			multiples[n] = append(multiples[n], divisor)
		}
	}

	for m := range multiples {
		sum += m
	}

	return
}
