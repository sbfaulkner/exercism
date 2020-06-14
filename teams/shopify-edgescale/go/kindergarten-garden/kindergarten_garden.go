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
type Garden map[string][]string

func sortChildren(children []string) ([]string, error) {
	unique := map[string]bool{}

	for _, c := range children {
		if unique[c] {
			return nil, errors.New("duplicate child name")
		}
		unique[c] = true
	}

	if sort.StringsAreSorted(children) {
		return children, nil
	}

	sorted := make([]string, len(children))
	copy(sorted, children)
	sort.Strings(sorted)

	return sorted, nil
}

// NewGarden instantiates a Garden given a diagram and a list of kindergarten children
func NewGarden(diagram string, children []string) (*Garden, error) {
	sortedChildren, err := sortChildren(children)
	if err != nil {
		return nil, err
	}

	rows := strings.Split(diagram, "\n")

	if len(rows) != 3 || rows[0] != "" || len(rows[1]) != 2*len(sortedChildren) || len(rows[1]) != len(rows[2]) {
		return nil, errors.New("wrong diagram format")
	}

	cups := rows[1:]

	g := Garden{}

	for i, c := range sortedChildren {
		low := i * 2
		high := low + 2
		g[c] = make([]string, 0, 4)
		for _, p := range cups[0][low:high] + cups[1][low:high] {
			name, ok := plantNames[p]
			if !ok {
				return nil, errors.New("invalid cup code")
			}
			g[c] = append(g[c], name)
		}
	}

	return &g, nil
}

// Plants returns the list of plants for a given child
func (g *Garden) Plants(child string) ([]string, bool) {
	p, ok := (*g)[child]
	return p, ok
}
