package pov

import (
	"errors"
	"fmt"
)

const testVersion = 2

// Graph is a directed graph
type Graph struct {
	root  *node
	nodes map[string]*node
}

// node is a node within a directed graph
type node struct {
	label    string
	parent   *node
	children []*node
}

// New creates a new (empty) directed graph
func New() *Graph {
	return &Graph{nodes: map[string]*node{}}
}

// AddNode adds a node with the specified label to the graph
func (g *Graph) AddNode(nodeLabel string) {
	if g.nodes[nodeLabel] != nil {
		panic(errors.New("duplicate node label"))
	}

	g.addNode(nodeLabel)
}

// AddArc creates an arc between the two specified nodes (adding the nodes if necessary)
func (g *Graph) AddArc(from, to string) {
	f := g.nodes[from]
	if f == nil {
		f = g.addNode(from)
	}

	t := g.nodes[to]
	if t == nil {
		t = g.addNode(to)
	}

	f.children = append(f.children, t)
	t.parent = f

	if g.root.label == to {
		g.root = f
	}
}

// ArcList returns a string representation of all arcs in the directed graph
func (g *Graph) ArcList() []string {
	if g.root == nil {
		return []string{}
	}

	return g.root.arcList()
}

// ChangeRoot returns a new directed graph with the specified (new) root
func (g *Graph) ChangeRoot(oldRoot, newRoot string) *Graph {
	ng := New()

	nr := g.nodes[newRoot]

	ng.copy(nr)
	ng.invert(nr)

	return ng
}

// addNode adds a node with the specifie label to the graph
func (g *Graph) addNode(nodeLabel string) *node {
	n := &node{label: nodeLabel}

	g.nodes[nodeLabel] = n

	if g.root == nil {
		g.root = n
	}

	return n
}

// copy recursively copies a node and its descendants to the graph
func (g *Graph) copy(n *node) {
	for _, c := range n.children {
		if g.nodes[c.label] != nil {
			continue
		}

		g.AddArc(n.label, c.label)
		g.copy(c)
	}
}

// invert inverts the relationship between a node and its parent
func (g *Graph) invert(n *node) {
	if n.parent == nil {
		return
	}

	g.AddArc(n.label, n.parent.label)
	g.copy(n.parent)
	g.invert(n.parent)
}

// arcList recursively generates string representations of the arcs starting at the current node
func (n *node) arcList() []string {
	a := []string{}

	for _, c := range n.children {
		a = append(a, fmt.Sprintf("%s -> %s", n.label, c.label))
		a = append(a, c.arcList()...)
	}

	return a
}
