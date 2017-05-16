package pythagorean

import "math"

const testVersion = 1

// Triplet represents a Pythagorean triplet
type Triplet [3]int

// Range returns a list of all Pythagorean triplets with sides in the
// range min to max inclusive.
func Range(min, max int) []Triplet {
	triplets := []Triplet{}

	for a := min; a <= max; a++ {
		for b := a; b <= max; b++ {
			c := math.Sqrt(float64(a*a + b*b))

			if math.Trunc(c) == c {
				if int(c) > max {
					break
				}

				triplets = append(triplets, Triplet{a, b, int(c)})
			}
		}
	}

	return triplets
}

// Sum returns a list of all Pythagorean triplets where the sum a+b+c
// (the perimeter) is equal to p.
func Sum(p int) []Triplet {
	triplets := []Triplet{}

	for _, triplet := range Range(1, p/2) {
		if triplet[0]+triplet[1]+triplet[2] == p {
			triplets = append(triplets, triplet)
		}
	}

	return triplets
}
