package kindergarten

import (
	"errors"
	"sort"
	"strings"
)

const testVersion = 1

// Garden holds the data for the classroom garden and the names of the students
type Garden struct {
	plants   []string
	children []string
}

var plantNames = map[rune]string{
	'C': "clover",
	'G': "grass",
	'R': "radishes",
	'V': "violets",
}

// standard errors for Garden
var (
	ErrInvalidGarden  = errors.New("garden: invalid diagram")
	ErrDuplicateChild = errors.New("garden: duplicate child")
)

// NewGarden constructs a new Garden
func NewGarden(diagram string, children []string) (*Garden, error) {
	g := Garden{}

	plants := strings.Split(diagram, "\n")[1:]
	if len(plants) != 2 {
		return nil, ErrInvalidGarden
	}

	for _, p := range plants {
		if len(p) != 2*len(children) {
			return nil, ErrInvalidGarden
		}

		if strings.IndexFunc(p, func(r rune) bool { _, ok := plantNames[r]; return !ok }) >= 0 {
			return nil, ErrInvalidGarden
		}
	}

	g.plants = plants

	for _, c := range children {
		i := sort.SearchStrings(g.children, c)
		if i < len(g.children) && g.children[i] == c {
			return nil, ErrDuplicateChild
		}

		g.children = append(g.children[0:i], append([]string{c}, g.children[i:]...)...)
	}

	return &g, nil
}

// Plants returns the plants for a given child
func (g *Garden) Plants(child string) ([]string, bool) {
	c := sort.SearchStrings(g.children, child)
	if c == len(g.children) {
		return []string{}, false
	}

	p := make([]string, 4)

	for irow, row := range g.plants {
		for icol, col := range row[c*2 : c*2+2] {
			p[irow*2+icol] = plantNames[col]
		}
	}

	return p, true
}
