package triangle

import "math"

const testVersion = 3

// Kind defines the types of triangles
type Kind int

// Pick values for the following identifiers used by the test program.
const (
	NaT Kind = 0 // not a triangle
	Equ Kind = 1 // equilateral
	Iso Kind = 2 // isosceles
	Sca Kind = 3 // scalene
)

func (k Kind) String() (text string) {
	switch k {
	case NaT:
		text = "NaT"
	case Equ:
		text = "Equ"
	case Iso:
		text = "Iso"
	case Sca:
		text = "Sca"
	}

	return
}

func isSide(f float64) bool {
	return f > 0 && !math.IsNaN(f) && !math.IsInf(f, 0)
}

func isTriangle(a, b, c float64) bool {
	if !isSide(a) || !isSide(b) || !isSide(c) {
		return false
	}

	if a+b < c || b+c < a || c+a < b {
		return false
	}

	return true
}

// KindFromSides identifies the kind of triangle given its side lengths
func KindFromSides(a, b, c float64) (kind Kind) {

	if !isTriangle(a, b, c) {
		kind = NaT
	} else if a == b && b == c {
		kind = Equ
	} else if a == b || b == c || c == a {
		kind = Iso
	} else {
		kind = Sca
	}

	return
}
