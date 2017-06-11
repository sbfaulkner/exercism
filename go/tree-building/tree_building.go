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

	sort.Slice(records, func(i, j int) bool { return records[i].ID < records[j].ID })

	nodes := map[int]*Node{}

	for id, r := range records {
		if r.ID > id {
			return nil, fmt.Errorf("tree: non-contiguous record ID - %d", r.ID)
		}

		n := &Node{ID: r.ID}
		nodes[r.ID] = n

		if r.ID == 0 {
			if r.Parent > 0 {
				return nil, errors.New("tree: parent specified for root node")
			}
		} else {
			if r.ID == r.Parent {
				return nil, fmt.Errorf("tree: self-referential node")
			}

			if nodes[r.Parent] == nil {
				return nil, fmt.Errorf("tree: missing parent - %d", r.Parent)
			}

			nodes[r.Parent].addChild(n)
		}
	}

	return nodes[0], nil
}

func (n *Node) addChild(child *Node) {
	n.Children = append(n.Children, child)
}
