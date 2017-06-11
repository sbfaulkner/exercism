package kindergarten

const testVersion = 1

// Garden holds the data for the classroom garden and the names of the students
type Garden struct {
	rows     []string
	children []string
}

// NewGarden constructs a new Garden
func NewGarden(diagram string, children []string) (*Garden, error) {
	return nil, nil
}

// Plants returns the plants for a given child
func (g *Garden) Plants(child string) ([]string, bool) {
	return []string{}, false
}
