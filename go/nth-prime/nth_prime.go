package prime

const testVersion = 1

// knownPrimes stores the known primes as they are found.
var knownPrimes = []int{2, 3}

// Nth returns the nth prime number
func Nth(n int) (int, bool) {
	if n < 1 {
		return 0, false
	}

	for n > len(knownPrimes) {
		findPrime()
	}

	return knownPrimes[n-1], true
}

// findPrime adds another prime to the collection of known primes.
func findPrime() {
	n := knownPrimes[len(knownPrimes)-1] + 2

	for {
		if !isDivisibleByKnownPrime(n) {
			break
		}

		n += 2
	}

	knownPrimes = append(knownPrimes, n)
}

// isDivisibleByKnownPrime checks whether or not a number is divisible by any of the known primes.
func isDivisibleByKnownPrime(n int) bool {
	for _, f := range knownPrimes {
		if n%f == 0 {
			return true
		}
	}

	return false
}
