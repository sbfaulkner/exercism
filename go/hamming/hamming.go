package hamming

import "errors"
import "fmt"

const testVersion = 6

// max returns the maximum of two int values
func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// Distance returns the Hamming distance between two strings representing DNA strands
func Distance(a, b string) (distance int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprint(r))
		}
	}()

	length := max(len(a), len(b))

	for i := 0; i < length; i++ {
		if a[i] != b[i] {
			distance++
		}
	}

	return
}
