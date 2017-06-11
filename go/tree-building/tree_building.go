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

type Mismatch struct{}

func (m Mismatch) Error() string {
	return "incorrect number of nodes"
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	root := &Node{}
	todo := []*Node{root}
	n := 1

	for {
		if len(todo) == 0 {
			break
		}

		newTodo := []*Node(nil)

		for _, t := range todo {
			for _, r := range records {
				if r.Parent == t.ID {
					if r.ID < t.ID {
						return nil, errors.New("parent exceeds id")
					}

					if r.ID == t.ID {
						if r.ID != 0 {
							return nil, fmt.Errorf("node references self")
						}
					} else {
						n++
						nn := &Node{ID: r.ID}
						newTodo = append(newTodo, nn)

						t.addChild(nn)
					}
				}
			}
		}
		todo = newTodo
	}

	if n != len(records) {
		return nil, Mismatch{}
	}

	if err := root.chk(len(records)); err != nil {
		return nil, err
	}

	return root, nil
}

func (n *Node) addChild(child *Node) {
	pos := sort.Search(len(n.Children), func(i int) bool { return n.Children[i].ID >= child.ID })
	n.Children = append(n.Children[0:pos], append([]*Node{child}, n.Children[pos:]...)...)
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
