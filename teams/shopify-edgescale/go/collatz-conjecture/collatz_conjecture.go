package collatzconjecture

import "errors"

// CollatzConjecture returns the number of steps to converge on 1
func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return -1, errors.New("n must be a positive integer")
	}

	s := 0

	for n > 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}

		s++
	}

	return s, nil
}
