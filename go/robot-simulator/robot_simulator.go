package robot

const testVersion = 3

// define directions
const (
	N Dir = iota
	E
	S
	W
)

// Right turns Step1Robot to the right
func Right() {
	Step1Robot.Dir = (Step1Robot.Dir + 1) % 4
}

// Left turns Step1Robot to the left
func Left() {
	Step1Robot.Dir = (Step1Robot.Dir + 3) % 4
}

// Advance moves Step1Robot one position forward in the direction it's facing
func Advance() {
	switch Step1Robot.Dir {
	case N:
		Step1Robot.Y++
	case S:
		Step1Robot.Y--
	case E:
		Step1Robot.X++
	case W:
		Step1Robot.X--
	}
}

// String converts a direction to a string representation
func (d Dir) String() (direction string) {
	switch d {
	case N:
		direction = "North"
	case S:
		direction = "South"
	case E:
		direction = "East"
	case W:
		direction = "West"
	}

	return
}
