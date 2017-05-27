package stringset

import "fmt"
import "strings"

const testVersion = 4

// Set is a collection of unique string values.
type Set map[string]bool

// New creates a new (empty) Set.
func New() Set {
	return map[string]bool{}
}

// NewFromSlice creates a new Set containing the strings in the provided slice.
func NewFromSlice(input []string) Set {
	s := New()

	for _, i := range input {
		s[i] = true
	}

	return s
}

// String returns a string representation of a Set.
func (s Set) String() string {
	output := []string{}

	for v := range s {
		output = append(output, fmt.Sprintf("%q", v))
	}

	return fmt.Sprintf("{%s}", strings.Join(output, ", "))
}

// IsEmpty returns true if the Set is empty.
func (s Set) IsEmpty() bool {
	return len(s) == 0
}

// Has returns true if the Set contains the provided string.
func (s Set) Has(v string) bool {
	return s[v]
}

// Subset returns true if s2 is a subset of s1.
func Subset(s1, s2 Set) bool {
	for v := range s1 {
		if !s2.Has(v) {
			return false
		}
	}

	return true
}

// Disjoint returns true if the provided sets are disjoint.
func Disjoint(s1, s2 Set) bool {
	for v := range s1 {
		if s2.Has(v) {
			return false
		}
	}
	return true
}

// Equal returns true if the provided sets are equal.
func Equal(s1, s2 Set) bool {
	return len(s1) == len(s2) && Subset(s1, s2)
}

// Add adds a string to the set.
func (s Set) Add(v string) {
	s[v] = true
}

// Intersection returns the intersection of two sets.
func Intersection(s1, s2 Set) Set {
	s := New()

	for v := range s1 {
		if s2.Has(v) {
			s[v] = true
		}
	}

	return s
}

// Difference returns the difference of two sets.
func Difference(s1, s2 Set) Set {
	for v := range s2 {
		delete(s1, v)
	}

	return s1
}

// Union returns the union of two sets.
func Union(s1, s2 Set) Set {
	for v := range s2 {
		s1[v] = true
	}

	return s1
}
