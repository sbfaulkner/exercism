package binarysearchtree

const testVersion = 1

// SearchTreeData represents a binary search tree
type SearchTreeData struct {
	left  *SearchTreeData
	data  int
	right *SearchTreeData
}

// Bst returns a new node for a binary search tree
func Bst(n int) *SearchTreeData {
	return &SearchTreeData{data: n}
}

// Insert adds a value into a binary tree
func (s *SearchTreeData) Insert(n int) {
	if n <= s.data {
		if s.left != nil {
			s.left.Insert(n)
		} else {
			s.left = Bst(n)
		}
	} else {
		if s.right != nil {
			s.right.Insert(n)
		} else {
			s.right = Bst(n)
		}
	}
}

// MapString maps the tree data using a provided function to return an array of string values
func (s *SearchTreeData) MapString(f func(int) string) (result []string) {
	if s.left != nil {
		result = append(s.left.MapString(f), f(s.data))
	} else {
		result = []string{f(s.data)}
	}

	if s.right != nil {
		result = append(result, s.right.MapString(f)...)
	}

	return
}

// MapInt maps the tree data using a provided function to return an array of int values
func (s *SearchTreeData) MapInt(f func(int) int) (result []int) {
	if s.left != nil {
		result = append(s.left.MapInt(f), f(s.data))
	} else {
		result = []int{f(s.data)}
	}

	if s.right != nil {
		result = append(result, s.right.MapInt(f)...)
	}

	return
}
