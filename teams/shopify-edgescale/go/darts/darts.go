package darts

import "math"

// Score returns the earned points in a single toss of a Darts game
func Score(x, y float64) int {
	r := math.Sqrt(x*x + y*y)

	switch {
	case r <= 1:
		return 10
	case r <= 5:
		return 5
	case r <= 10:
		return 1
	}

	return 0
}
