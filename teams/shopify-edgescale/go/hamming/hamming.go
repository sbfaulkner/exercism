package hamming

import "errors"

// Distance calculates the Hamming Distance between two DNA strands
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("DNA strands must be the same length")
	}

	d := 0

	for i, aa := range a {
		if aa != []rune(b)[i] {
			d++
		}
	}

	return d, nil
}
