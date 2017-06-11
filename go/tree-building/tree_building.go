package tree

import (
	"errors"
	"fmt"
	"sort"
)

const testVersion = 4

type Record struct {
	ID, Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	nodes := map[int]*Node{}

	for _, r := range records {
		if r.ID == 0 && r.Parent > 0 {
			return nil, errors.New("tree: root can't have parent")
		}

		n := nodes[r.ID]

		if n == nil {
			nodes[r.ID] = &Node{ID: r.ID}
			n = nodes[r.ID]
		} else {
			n.ID = r.ID
		}

		p := nodes[r.Parent]

		if p == nil {
			nodes[r.Parent] = &Node{ID: -1}
			p = nodes[r.Parent]
		}

		if n.ID > 0 {
			if n.ID == p.ID {
				return nil, fmt.Errorf("tree: self referential node - %d", n.ID)
			}

			if n.ID < p.ID {
				return nil, fmt.Errorf("tree: invalid parent node - %d < %d", n.ID, p.ID)
			}

			if n.isAncestorOf(p) {
				return nil, fmt.Errorf("tree: circular reference - %d is ancestor of %d", n.ID, p.ID)
			}

			p.addChild(n)
		}
	}

	for id, n := range nodes {
		if n.ID == -1 {
			return nil, fmt.Errorf("tree: missing node - %d", id)
		}

		if id >= len(nodes) {
			return nil, fmt.Errorf("tree: invalid node id - %d", id)
		}
	}

	return nodes[0], nil
}

func (n *Node) addChild(child *Node) {
	pos := sort.Search(len(n.Children), func(i int) bool { return n.Children[i].ID >= child.ID })
	n.Children = append(n.Children[0:pos], append([]*Node{child}, n.Children[pos:]...)...)
}

func (n *Node) isAncestorOf(child *Node) bool {
	pos := sort.Search(len(n.Children), func(i int) bool { return n.Children[i].ID >= child.ID })

	if pos < len(n.Children) && n.Children[pos].ID == child.ID {
		return true
	}

	for _, c := range n.Children {
		if c.isAncestorOf(child) {
			return true
		}
	}

	return false
}

func (n *Node) chk(m int) error {
	if n.ID > m {
		return fmt.Errorf("z")
	}

	if n.ID == m {
		return fmt.Errorf("y")
	}

	for _, c := range n.Children {
		if err := c.chk(m); err != nil {
			return err
		}
	}

	return nil
}
