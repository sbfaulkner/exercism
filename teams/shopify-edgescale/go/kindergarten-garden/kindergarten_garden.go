package kindergarten

import (
	"errors"
	"sort"
	"strings"
)

var plantNames = map[rune]string{
	'G': "grass",
	'C': "clover",
	'R': "radishes",
	'V': "violets",
}

// Garden is an instance of a kindergarten class's garden
type Garden struct {
	cups     [][]rune
	children []string
}

func parseChildren(children []string) ([]string, error) {
	unique := map[string]bool{}

	for _, c := range children {
		if _, ok := unique[c]; ok {
			return nil, errors.New("duplicate name")
		}
		unique[c] = true
	}

	sorted := make([]string, len(children))
	copy(sorted, children)
	sort.Strings(sorted)

	return sorted, nil
}

func parseCups(diagram string) ([][]rune, error) {
	rows := strings.Split(diagram, "\n")
	if rows[0] != "" {
		return nil, errors.New("wrong diagram format")
	}

	if len(rows[1]) != len(rows[2]) {
		return nil, errors.New("mismatched rows")
	}

	if len(rows[1])%2 != 0 {
		return nil, errors.New("odd number of cups")
	}

	for _, r := range rows[1:3] {
		for _, c := range r {
			if _, ok := plantNames[c]; !ok {
				return nil, errors.New("invalid cup code")
			}
		}
	}
	return [][]rune{[]rune(rows[1]), []rune(rows[2])}, nil
}

// NewGarden instantiates a Garden given a diagram and a list of kindergarten children
func NewGarden(diagram string, children []string) (*Garden, error) {
	parsedCups, err := parseCups(diagram)
	if err != nil {
		return nil, err
	}

	parsedChildren, err := parseChildren(children)
	if err != nil {
		return nil, err
	}

	g := Garden{
		cups:     parsedCups,
		children: parsedChildren,
	}

	return &g, nil
}

// Plants returns the list of plants for a given child
func (g *Garden) Plants(child string) ([]string, bool) {
	i := sort.SearchStrings(g.children, child)

	if i == len(g.children) {
		return nil, false
	}

	plants := []string{
		plantNames[g.cups[0][i*2]],
		plantNames[g.cups[0][i*2+1]],
		plantNames[g.cups[1][i*2]],
		plantNames[g.cups[1][i*2+1]],
	}

	return plants, true
}
