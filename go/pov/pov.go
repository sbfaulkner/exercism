package pov

import (
	"fmt"
)

const testVersion = 2

// Graph is a directed graph
type Graph map[string]string

// New creates a new (empty) directed graph
func New() *Graph {
	return &Graph{}
}

// AddNode adds a node with the specified label to the graph
func (g *Graph) AddNode(nodeLabel string) {
	(*g)[nodeLabel] = ""
}

// AddArc creates an arc between the two specified nodes (adding the nodes if necessary)
func (g *Graph) AddArc(from, to string) {
	(*g)[to] = from
}

// ArcList returns a string representation of all arcs in the directed graph
func (g *Graph) ArcList() []string {
	a := make([]string, 0, len(*g))

	for t, f := range *g {
		if f == "" {
			continue
		}

		a = append(a, fmt.Sprintf("%s -> %s", f, t))
	}

	return a
}

// ChangeRoot returns a new directed graph with the specified (new) root
func (g *Graph) ChangeRoot(oldRoot, newRoot string) *Graph {
	ng := make(Graph, len(*g))

	for t, f := range *g {
		ng[t] = f
	}

	ng.invert(newRoot)

	delete(ng, newRoot)

	return &ng
}

func (g *Graph) invert(label string) {
	f := (*g)[label]

	if f != "" {
		g.invert(f)
		(*g)[f] = label
	}
}
