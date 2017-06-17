package pov

import (
	"errors"
	"fmt"
)

const testVersion = 2

type node struct {
	label    string
	parent   *node
	children []*node
}

type Graph struct {
	root  *node
	nodes map[string]*node
}

func New() *Graph {
	return &Graph{nodes: map[string]*node{}}
}

func (g *Graph) AddNode(nodeLabel string) {
	if g.nodes[nodeLabel] != nil {
		panic(errors.New("duplicate node label"))
	}

	g.addNode(nodeLabel)
}

func (g *Graph) addNode(nodeLabel string) *node {
	n := &node{label: nodeLabel}

	g.nodes[nodeLabel] = n

	if g.root == nil {
		g.root = n
	}

	return n
}

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

func (g *Graph) ArcList() []string {
	if g.root == nil {
		return []string{}
	}

	return g.root.arcList()
}

func (n *node) arcList() []string {
	a := []string{}

	for _, c := range n.children {
		a = append(a, fmt.Sprintf("%s -> %s", n.label, c.label))
		a = append(a, c.arcList()...)
	}

	return a
}

func (g *Graph) ChangeRoot(oldRoot, newRoot string) *Graph {
	ng := New()

	nr := g.nodes[newRoot]

	ng.copy(nr)
	ng.invert(nr)

	return ng
}

func (g *Graph) copy(n *node) {
	for _, c := range n.children {
		if g.nodes[c.label] != nil {
			continue
		}

		g.AddArc(n.label, c.label)
		g.copy(c)
	}
}

func (g *Graph) invert(n *node) {
	if n.parent == nil {
		return
	}

	g.AddArc(n.label, n.parent.label)
	g.copy(n.parent)
	g.invert(n.parent)
}
