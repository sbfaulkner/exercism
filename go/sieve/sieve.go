package sieve

const testVersion = 1

// Sieve uses the Sieve of Eratosthenes to find all the primes from 2 up to a given number.
func Sieve(limit int) []int {
	marks := make(map[int]bool, limit-1)

	for i := 2; i <= limit; i++ {
		if marks[i] {
			continue
		}

		for n := i * 2; n <= limit; n += i {
			marks[n] = true
		}
	}

	primes := make([]int, limit-1-len(marks))
	count := 0

	for i := 2; i <= limit; i++ {
		if !marks[i] {
			primes[count] = i
			count++
		}
	}

	return primes
}
