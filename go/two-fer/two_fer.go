// Package twofer implements a silly sharing descriptor.
package twofer

import "fmt"

// ShareWith returns a string describing sharing with the specified person.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return fmt.Sprintf("One for %v, one for me.", name)
}
